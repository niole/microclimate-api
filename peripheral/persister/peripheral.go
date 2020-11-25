package persister

import (
	"api/database"
	api "api/protobuf"
	"context"
	"database/sql"
	"log"
)

func initTable(ctx context.Context, pool *sql.DB) error {
	_, peripheralTableCreateErr := pool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Peripherals (
            Id varchar(36) PRIMARY KEY NOT NULL,
            OwnerUserId varchar(36) NOT NULL,
            DeploymentId varchar(36) NOT NULL,
            HardwareId varchar(36) NOT NULL,
            Type ENUM('THERMAL', 'PARTICLE') NOT NULL,
            Unit ENUM('F', 'C', 'PM2.5', '%') NOT NULL,
            Name varchar(255) NOT NULL
        );`,
	)
	if peripheralTableCreateErr != nil {
		log.Fatalf("Failed to create peripherals table. error: %v", peripheralTableCreateErr)
	}
	return peripheralTableCreateErr
}

func CreatePeripheral(
	ctx context.Context,
	ownerId string,
	deploymentId string,
	hardwareId string,
	pType api.NewPeripheral_PeripheralType,
	unit string,
	name string,
) (*api.Peripheral, error) {
	db := database.Get(initTable)

	_, err := db.ExecContext(
		ctx,
		`INSERT INTO Peripherals
        (Id, OwnerUserId, DeploymentId, HardwareId, Type, Unit, Name) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		&hardwareId,
		&ownerId,
		&deploymentId,
		&hardwareId, // TODO get rid of hardware id
		api.Peripheral_PeripheralType_name[int32(pType)],
		&unit,
		&name,
	)

	if err != nil {
		log.Printf(
			"Failed to insert new peripheral with hardware id: %v, for user %v, err: %v",
			&hardwareId,
			&ownerId,
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
		&ownerId,
		&deploymentId,
		&hardwareId,
	).Scan, &peripheral)

	return &peripheral, err
}

func RemovePeripheral(ctx context.Context, peripheral *api.Peripheral) error {
	db := database.Get(initTable)

	_, err := db.ExecContext(
		ctx,
		`DELETE FROM Peripherals
        WHERE ID = ? `,
		&peripheral.Id,
	)

	return err
}

func GetDeploymentPeripherals(ctx context.Context, deploymentId string) ([]api.Peripheral, error) {
	db := database.Get(initTable)

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

func GetPeripheralById(ctx context.Context, peripheralId string) (*api.Peripheral, error) {
	db := database.Get(initTable)

	var peripheral api.Peripheral

	err := ScanOnePeripheral(db.QueryRowContext(
		ctx,
		`SELECT * FROM Peripherals
        WHERE Id = ?
        LIMIT 1;`,
		&peripheralId,
	).Scan, &peripheral)

	return &peripheral, err
}
