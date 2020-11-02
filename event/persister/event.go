package persister

import (
	"api/database"
	api "api/protobuf"
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	"log"
	"strings"
	"time"
)

func SaveEvent(ctx context.Context, event *api.NewMeasurementEvent) error {
	db := database.Get(initTable)

	time, _ := ConvertPBTimeToTime(event.TimeStamp)

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
	db := database.Get(initTable)

	in := strings.Join(request.PeripheralIds, ",")

	endtime, _ := ConvertPBTimeToTime(request.EndTime)
	starttime, _ := ConvertPBTimeToTime(request.StartTime)

	events, err := ScanEvents(db.QueryContext(
		ctx,
		`SELECT * FROM PeripheralEvents WHERE
		Timestamp BETWEEN ? AND ?
		AND DeploymentId = ?
		AND PeripheralId IN (?);`,
		starttime,
		endtime,
		&request.DeploymentId,
		&in,
	))

	if err != nil {
		log.Printf("Failed to get events out of db, error %v", err)
	}

	return events, err
}

func initTable(ctx context.Context, pool *sql.DB) error {
	_, peripheralEventsTableCreateErr := pool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS PeripheralEvents (
                Id varchar(36) PRIMARY KEY NOT NULL,
				PeripheralId varchar(36) NOT NULL,
                DeploymentId varchar(36) NOT NULL,
                Value smallint NOT NULL,
				Timestamp timestamp NOT NULL
            );`,
	)
	if peripheralEventsTableCreateErr != nil {
		log.Fatalf(
			"Failed to create peripheral events table. error: %v",
			peripheralEventsTableCreateErr,
		)
	}
	return peripheralEventsTableCreateErr
}

func ConvertPBTimeToTime(pbtime *timestamppb.Timestamp) (time.Time, error) {
	time, err := ptypes.Timestamp(pbtime)
	if err != nil {
		log.Fatalf("Failed to convert timestamp to sql consumable type, error %v", err)
	}

	return time, err
}