package event

import (
	api "api/protobuf"
	connection "api/server/impl/database"
	"context"
	"github.com/golang/protobuf/ptypes"
	"log"
)

func SaveEvent(ctx context.Context, event *api.NewMeasurementEvent) error {
	db := connection.GetConnectionPool()

	time, conversionErr := ptypes.Timestamp(event.TimeStamp)
	if conversionErr != nil {
		log.Fatalf("Failed to convert timestamp to sql consumable type, erorr %v", conversionErr)
	}

	_, err := db.ExecContext(
		ctx,
		`INSERT INTO PeripheralEvents
		(Id, PeripheralId, DeploymentId, Value, Timestamp)
		VALUES (UUID(), ?, ?, ?, ?);`,
		&event.PeripheralId,
		&event.DeploymentId,
		&event.Value,
		time,
	)

	return err
}

func FilterEvents(ctx context.Context, request *api.MeasurementEventFilterRequest) ([]api.MeasurementEvent, error) {
	db := connection.GetConnectionPool()

	events, err := ScanEvents(db.QueryContext(
		ctx,
		`SELECT * FROM PeripheralEvents;`,
	))

	if err != nil {
		log.Printf("Failed to get events out of db, error %v", err)
	}

	return events, err
}
