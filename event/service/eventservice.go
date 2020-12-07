package service

import (
	"api/clients"
	"api/event/persister"
	api "api/protobuf"
	"context"
	"log"
	"time"
)

type PeripheralEventService struct {
	api.UnimplementedPeripheralMeasurementEventServiceServer
}

func (s PeripheralEventService) SendEvent(ctx context.Context, in *api.NewMeasurementEvent) (*api.Empty, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Printf("Received event from peripheral %v, value %v", in.PeripheralId, in.Value)

	conn, err := clients.PeripheralClientConnection()
	defer conn.Close()
	client := api.NewPeripheralManagementServiceClient(conn)
	p, err := client.GetPeripheral(ctx, &api.GetPeripheralRequest{
		PeripheralId: in.PeripheralId,
	})

	if err == nil && p != nil && p.Id == in.PeripheralId {
		err = persister.SaveEvent(cancellableCtx, in)

		if err != nil {
			log.Printf("Failed to save event %v, error %v", in, err)
		} else {
			log.Printf("Saved new event for peripheral with id %v", in.PeripheralId)
		}
	} else {
		log.Printf("Failed to verify inputs %v", err)
	}

	return &api.Empty{}, err
}

func (s PeripheralEventService) FilterEvents(
	request *api.MeasurementEventFilterRequest,
	stream api.PeripheralMeasurementEventService_FilterEventsServer,
) error {
	cancellableCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	events, err := persister.FilterEvents(cancellableCtx, request)
	if err != nil {
		log.Printf("Failed to get requested events %v, error %v", request, err)
	} else {
		log.Printf("Got %v events for peripheral with id %v", len(events), request.PeripheralId)

		for _, s := range events {
			err = stream.Send(&s)
			if err != nil {
				log.Printf("Failed to element over stream, error %v", err)
			}
		}
	}

	return err
}

func (s PeripheralEventService) DeletePeripheralEvents(ctx context.Context, in *api.DeletePeripheralEventsRequest) (*api.Empty, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Printf("Deleting events for peripheral %v", in.PeripheralId)

	err := persister.DeletePeripheralEvents(cancellableCtx, in.PeripheralId)

	if err != nil {
		log.Printf("Failed to remove events for peripheral %v, error %v", in.PeripheralId, err)
	} else {
		log.Printf("Deleted events for peripheral %v", in.PeripheralId)
	}

	return &api.Empty{}, err
}

func (s PeripheralEventService) MostRecentDeploymentEvents(
	in *api.MostRecentEventsForDeploymentRequest,
	stream api.PeripheralMeasurementEventService_MostRecentDeploymentEventsServer,
) error {
	cancellableCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("Getting most recent events for deployment %v", in.DeploymentId)

	events, err := persister.MostRecentDeploymentEvents(cancellableCtx, in.DeploymentId)
	if err != nil {
		log.Printf("Failed to get most recent events for deployment requested events %v, error %v", in.DeploymentId, err)
	} else {
		log.Printf("Got %v recent events for peripherals in deployment with id %v", len(events), in.DeploymentId)

		for _, s := range events {
			err = stream.Send(&s)
			if err != nil {
				log.Printf("Failed to send element over stream, error %v", err)
			}
		}
	}

	return err
}
