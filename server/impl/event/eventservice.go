package event

import (
	api "api/protobuf"
	"api/server/impl/event/persister"
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

	saveError := persister.SaveEvent(cancellableCtx, in)

	if saveError != nil {
		log.Fatalf("Failed to save event %v, error %v", in, saveError)
	}

	return &api.Empty{}, saveError
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
