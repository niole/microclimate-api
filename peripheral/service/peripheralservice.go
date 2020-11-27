package service

import (
	"api/clients"
	"api/peripheral/persister"
	api "api/protobuf"
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
	peripheral, err := persister.CreatePeripheral(
		cancellableCtx,
		newPeripheral.OwnerUserId,
		newPeripheral.DeploymentId,
		newPeripheral.HardwareId,
		newPeripheral.Type,
		newPeripheral.Unit,
		newPeripheral.Name,
	)
	if err != nil {
		log.Printf("Failed to create peripheral: %v, err: %v", newPeripheral, err)
	}
	return peripheral, err
}

func (s PeripheralManagementService) RemovePeripheral(ctx context.Context, peripheral *api.Peripheral) (*api.Empty, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := persister.RemovePeripheral(cancellableCtx, peripheral)

	if err == nil {
		log.Printf("Removing events for peripheral %v", peripheral.Id)
		conn, err := clients.EventsClientConnection()
		eventsClient := api.NewPeripheralMeasurementEventServiceClient(conn)
		_, err = eventsClient.DeletePeripheralEvents(cancellableCtx, &api.DeletePeripheralEventsRequest{PeripheralId: peripheral.Id})
		if err != nil {
			log.Printf("Failed to delete events for peripheral %v", peripheral.Id)
		}
	}

	return &api.Empty{}, err
}

func (s PeripheralManagementService) GetDeploymentPeripherals(
	request *api.GetDeploymentPeripheralsRequest,
	peripheralStream api.PeripheralManagementService_GetDeploymentPeripheralsServer,
) error {
	cancellableCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Getting deployment %v peripherals", request.DeploymentId)

	peripherals, err := persister.GetDeploymentPeripherals(cancellableCtx, request.DeploymentId)
	if err != nil {
		log.Printf("Failed to get deployment peripherals, error %v", err)
	} else {
		for _, p := range peripherals {
			err = peripheralStream.Send(&p)
			if err != nil {
				log.Printf("Failed to send peripheral over stream %v", p)
			}
		}
	}

	return err
}

func (s PeripheralManagementService) GetPeripheral(ctx context.Context, request *api.GetPeripheralRequest) (*api.Peripheral, error) {

	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	peripheral, err := persister.GetPeripheralById(cancellableCtx, request.PeripheralId)

	if err != nil {
		log.Printf("Failed to get peripheral %v error %v", request, err)
	}

	return peripheral, err
}
