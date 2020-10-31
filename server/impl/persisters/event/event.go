package event

import (
	api "api/protobuf"
	connection "api/server/impl/database"
	"context"
	"github.com/golang/protobuf/ptypes"
	"log"
)

func SaveEvent(ctx context.Context, event *api.MeasurementEvent) error {
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
