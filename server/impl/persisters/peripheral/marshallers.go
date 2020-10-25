package peripheral

import (
	api "api/protobuf"
	sql "database/sql"
)

func ScanOnePeripheral(result *sql.Row, saveToPeripheral *api.Peripheral) error {
	var id string
	var ownerUserId string
	var deploymentId string
	var peripheralType string
	var hardwareId string
	err := result.Scan(&id, &ownerUserId, &deploymentId, &hardwareId, &peripheralType)

	*saveToPeripheral = api.Peripheral{
		OwnerUserId:  ownerUserId,
		DeploymentId: deploymentId,
		Id:           id,
		HardwareId:   hardwareId,
		Type:         api.Peripheral_PeripheralType(api.Peripheral_PeripheralType_value[peripheralType]),
	}

	return err
}
