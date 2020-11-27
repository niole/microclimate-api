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
	cancellableCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	events, err := persister.FilterEvents(cancellableCtx, request)
	if err != nil {
		log.Printf("Failed to get requested events %v, error %v", request, err)
	} else {
		for _, s := range events {
			err = stream.Send(&s)
			if err != nil {
				log.Printf("Failed to element over stream, error %v", err)
			}
		}
	}

	return err
}
