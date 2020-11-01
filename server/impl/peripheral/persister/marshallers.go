package persister

import (
	api "api/protobuf"
	sql "database/sql"
	"log"
)

type Scannable = func(dest ...interface{}) error

func ScanOnePeripheral(result Scannable, saveToPeripheral *api.Peripheral) error {
	var id string
	var ownerUserId string
	var deploymentId string
	var peripheralType string
	var hardwareId string
	err := result(&id, &ownerUserId, &deploymentId, &hardwareId, &peripheralType)

	*saveToPeripheral = api.Peripheral{
		OwnerUserId:  ownerUserId,
		DeploymentId: deploymentId,
		Id:           id,
		HardwareId:   hardwareId,
		Type:         api.Peripheral_PeripheralType(api.Peripheral_PeripheralType_value[peripheralType]),
	}

	return err
}

func ScanPeripherals(rows *sql.Rows, err error) ([]api.Peripheral, error) {
	defer rows.Close()

	peripherals := make([]api.Peripheral, 0)
	for rows.Next() {
		var peripheral api.Peripheral
		if err = ScanOnePeripheral(rows.Scan, &peripheral); err != nil {
			log.Fatal(err)
		}
		peripherals = append(peripherals, peripheral)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return peripherals, err
}
