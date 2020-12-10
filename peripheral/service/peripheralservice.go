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

	var hardwareId *string
	foundHardwareId := newPeripheral.HardwareId.GetData()
	if foundHardwareId != "" {
		hardwareId = &foundHardwareId
	}

	peripheral, err := persister.CreatePeripheral(
		cancellableCtx,
		newPeripheral.OwnerUserId,
		newPeripheral.DeploymentId,
		hardwareId,
		newPeripheral.Type,
		newPeripheral.Unit,
		newPeripheral.Name,
	)
	if err != nil {
		log.Printf("Failed to create peripheral: %v, err: %v", newPeripheral, err)
	} else {
		log.Printf("Successfully created new peripheral with id %v", peripheral.Id)
	}
	return peripheral, err
}

func (s PeripheralManagementService) RemovePeripheral(ctx context.Context, peripheral *api.Peripheral) (*api.Empty, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := persister.RemovePeripheral(cancellableCtx, peripheral)

	if err == nil {
		conn, err := clients.EventsClientConnection()
		eventsClient := api.NewPeripheralMeasurementEventServiceClient(conn)
		_, err = eventsClient.DeletePeripheralEvents(cancellableCtx, &api.DeletePeripheralEventsRequest{PeripheralId: peripheral.Id})
		if err != nil {
			log.Printf("Failed to delete events for peripheral %v", peripheral.Id)
		} else {
			log.Printf("Successfully removed peripheral with id %v", peripheral.Id)
		}
	} else {
		log.Printf("Failed to remove peripheral %v, err %v", peripheral.Id, err)
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
		log.Printf("Found total peripherals %v for deployment with id %v", len(peripherals), request.DeploymentId)

		for _, p := range peripherals {
			err = peripheralStream.Send(&p)
			if err != nil {
				log.Printf("Failed to send peripheral over stream %v", p)
			}
		}
	}

	return err
}

func (s PeripheralManagementService) LinkHardware(ctx context.Context, request *api.LinkHardwareRequest) (*api.Peripheral, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// do input validation
	foundPeripheral, err := s.GetPeripheral(ctx, &api.GetPeripheralRequest{PeripheralId: request.PeripheralId})
	if foundPeripheral != nil {

		peripheral, err := persister.LinkHardware(cancellableCtx, request.PeripheralId, request.HardwareId)

		if err != nil {
			log.Printf("Failed to link hardware to  peripheral %v error %v", request, err)
		} else {
			log.Printf("Successfully linked peripheral with id %v to hardware id %v", request.PeripheralId, request.HardwareId)
		}

		return peripheral, err
	} else {
		// couldn't find this peripheral
		// TODO also do authorization eventually
		log.Printf(
			"Couldn't find peripheral, couldn't link hardare with id %v to peripheral id %v, err %v",
			request.HardwareId,
			request.PeripheralId,
			err,
		)
	}

	return foundPeripheral, err
}

func (s PeripheralManagementService) GetPeripheral(ctx context.Context, request *api.GetPeripheralRequest) (*api.Peripheral, error) {

	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	peripheral, err := persister.GetPeripheralById(cancellableCtx, request.PeripheralId)

	if err != nil {
		log.Printf("Failed to get peripheral %v error %v", request, err)
	} else {
		log.Printf("Successfully got peripheral with id %v", request.PeripheralId)
	}

	return peripheral, err
}

func (s PeripheralManagementService) EditPeripheral(ctx context.Context, request *api.EditPeripheralRequest) (*api.Peripheral, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Printf("Editing peripheral with id %v", request.PeripheralId)

	peripheral, err := persister.GetPeripheralById(cancellableCtx, request.PeripheralId)

	if peripheral != nil {
		var newName *string
		foundName := request.NewName.GetData()
		if foundName != "" {
			newName = &foundName
		}

		var newType *api.Peripheral_PeripheralType
		if x, ok := request.NewType.GetKind().(*api.NullablePType_Data); ok {
			newType = &x.Data
		}

		updatedPeripheral, err := persister.EditPeripheral(
			cancellableCtx,
			request.PeripheralId,
			newName,
			newType,
		)

		if err != nil {
			log.Printf("Failed to edit peripheral with id %v. error: %v", request.PeripheralId, err)
		}

		return updatedPeripheral, err
	} else {
		if err != nil {
			log.Printf("Couldn't get peripheral with id %v. error: %v", request.PeripheralId, err)
		} else {
			log.Printf("Peripheral with id %v does not exist", request.PeripheralId)
		}
	}

	if err != nil {
		log.Printf("Failed to get peripheral %v error %v", request, err)
	} else {
		log.Printf("Successfully got peripheral with id %v", request.PeripheralId)
	}

	return peripheral, err
}
