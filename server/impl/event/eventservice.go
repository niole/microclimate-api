package event

import (
	api "api/protobuf"
	persister "api/server/impl/persisters/event"
	"context"
	"log"
	"time"
)

type PeripheralEventService struct {
	api.UnimplementedPeripheralMeasurementEventServiceServer
}

func (s PeripheralEventService) SendEvent(ctx context.Context, in *api.MeasurementEvent) (*api.Empty, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	saveError := persister.SaveEvent(cancellableCtx, in)

	if saveError != nil {
		log.Fatalf("Failed to save event %v, error %v", in, saveError)
	}

	return &api.Empty{}, saveError
}

func (s PeripheralEventService) FilterEvents(*api.MeasurementEventFilterRequest, api.PeripheralMeasurementEventService_FilterEventsServer) error {
	return nil
}
