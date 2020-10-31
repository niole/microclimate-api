package event

import (
	api "api/protobuf"
	sql "database/sql"
	"github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	"log"
)

type Scannable = func(dest ...interface{}) error

func ScanOneEvent(scan Scannable, event *api.MeasurementEvent) error {
	var id string
	var peripheralId string
	var deploymentId string
	var value int32
	var timestamp *sql.NullTime
	var parsedTimestamp *timestamppb.Timestamp

	err := scan(&id, &peripheralId, &deploymentId, &value, &timestamp)

	if timestamp.Valid {
		parsedTimestamp, err = ptypes.TimestampProto(timestamp.Time)
	}

	if err != nil {
		log.Printf("Something went wrong while scanning an event, error %v", err)
	} else {
		*event = api.MeasurementEvent{
			Id:           id,
			PeripheralId: peripheralId,
			DeploymentId: deploymentId,
			Value:        value,
			TimeStamp:    parsedTimestamp,
		}
	}

	return err
}

func ScanEvents(rows *sql.Rows, err error) ([]api.MeasurementEvent, error) {
	defer rows.Close()

	events := make([]api.MeasurementEvent, 0)

	if err == nil {
		for rows.Next() {
			var event api.MeasurementEvent
			if err = ScanOneEvent(rows.Scan, &event); err == nil {
				events = append(events, event)
			}
		}

		if err := rows.Err(); err != nil {
			log.Printf("Failed so scan some or all events for query, error %v", err)
		}
	} else {
		log.Printf("Something went wrong while doing events query, err %v", err)
	}

	return events, err
}
