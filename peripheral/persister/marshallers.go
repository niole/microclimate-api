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
	var hardwareId sql.NullString
	var unit string
	var name string

	err := result(
		&id,
		&ownerUserId,
		&deploymentId,
		&hardwareId,
		&peripheralType,
		&unit,
		&name,
	)

	var nullableHardwareId = &api.NullableString{
		Kind: &api.NullableString_Null{},
	}

	if hardwareId.Valid {
		nullableHardwareId = &api.NullableString{
			Kind: &api.NullableString_Data{
				Data: hardwareId.String,
			},
		}
	}

	*saveToPeripheral = api.Peripheral{
		OwnerUserId:  ownerUserId,
		DeploymentId: deploymentId,
		Id:           id,
		HardwareId:   nullableHardwareId,
		Type:         api.Peripheral_PeripheralType(api.Peripheral_PeripheralType_value[peripheralType]),
		Unit:         unit,
		Name:         name,
	}

	return err
}

func ScanPeripherals(rows *sql.Rows, err error) ([]api.Peripheral, error) {
	defer rows.Close()

	peripherals := make([]api.Peripheral, 0)
	for rows.Next() {
		var peripheral api.Peripheral
		if err = ScanOnePeripheral(rows.Scan, &peripheral); err != nil {
			log.Print(err)
		}
		peripherals = append(peripherals, peripheral)
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
	}

	return peripherals, err
}
