package peripheral

import (
	api "api/protobuf"
	connector "api/server/impl/database"
	"context"
	"log"
)

func CreatePeripheral(ctx context.Context, newPeripheral *api.NewPeripheral) (*api.Peripheral, error) {
	db := connector.GetConnectionPool()

	_, err := db.ExecContext(
		ctx,
		`INSERT INTO Peripherals
		(Id, OwnerUserId, DeploymentId, HardwareId, Type) VALUES (UUID(), ?, ?, ?, ?)`,
		&newPeripheral.OwnerUserId,
		&newPeripheral.DeploymentId,
		&newPeripheral.HardwareId,
		api.Peripheral_PeripheralType_name[int32(newPeripheral.Type)],
	)

	if err != nil {
		log.Printf(
			"Failed to insert new peripheral: %v, err: %v",
			&newPeripheral,
			err,
		)
		return nil, err
	}

	var peripheral api.Peripheral

	err = ScanOnePeripheral(db.QueryRowContext(
		ctx,
		`SELECT * FROM Peripherals
		WHERE OwnerUserId = ? AND DeploymentId = ? AND HardwareId = ?
		LIMIT 1;`,
		&newPeripheral.OwnerUserId,
		&newPeripheral.DeploymentId,
		&newPeripheral.HardwareId,
	).Scan, &peripheral)

	return &peripheral, err
}

func RemovePeripheral(ctx context.Context, peripheral *api.Peripheral) error {
	db := connector.GetConnectionPool()

	_, err := db.ExecContext(
		ctx,
		`DELETE FROM Peripherals
		WHERE ID = ? `,
		&peripheral.Id,
	)

	return err
}

func GetDeploymentPeripherals(ctx context.Context, deploymentId string) ([]api.Peripheral, error) {
	db := connector.GetConnectionPool()

	peripherals, err := ScanPeripherals(db.QueryContext(
		ctx,
		`SELECT * FROM Peripherals
		WHERE DeploymentId = ?;`,
		deploymentId,
	))

	if err != nil {
		log.Fatalf("Failed to get peripherals for deployment: %v, err: %v", deploymentId, err)
	}

	return peripherals, err
}
