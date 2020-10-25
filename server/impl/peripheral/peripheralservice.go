package peripheral

import (
	api "api/protobuf"
	persister "api/server/impl/persisters/peripheral"
	"context"
	"log"
	"time"
)

type PeripheralManagementService struct {
	api.UnimplementedPeripheralManagementServiceServer
}

func (s PeripheralManagementService) CreatePeripheral(ctx context.Context, newPeripheral *api.NewPeripheral) (*api.Peripheral, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	peripheral, err := persister.CreatePeripheral(cancellableCtx, newPeripheral)
	if err != nil {
		log.Printf("Failed to create peripheral: %v, err: %v", newPeripheral, err)
	}
	return peripheral, err
}

func (s PeripheralManagementService) RemovePeripheral(ctx context.Context, peripheral *api.Peripheral) (*api.Empty, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := persister.RemovePeripheral(cancellableCtx, peripheral)

	return &api.Empty{}, err
}
