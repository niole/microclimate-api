package persister

import (
	"api/database"
	api "api/protobuf"
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	"log"
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

	endtime, _ := ConvertPBTimeToTime(request.EndTime)
	starttime, _ := ConvertPBTimeToTime(request.StartTime)

	log.Printf(
		"Getting events. start time %v, end time %v, peripheral id %v, deploymentid %v",
		starttime,
		endtime,
		request.PeripheralId,
		request.DeploymentId,
	)

	events, err := ScanEvents(db.QueryContext(
		ctx,
		`SELECT * FROM PeripheralEvents WHERE
        Timestamp BETWEEN ? AND ?
        AND DeploymentId = ?
        AND PeripheralId = ?
		ORDER BY Timestamp ASC;`,
		starttime,
		endtime,
		&request.DeploymentId,
		&request.PeripheralId,
	))

	if err != nil {
		log.Printf("Failed to get events out of db, error %v", err)
	}

	log.Printf("Found %v events", len(events))

	return events, err
}

func DeletePeripheralEvents(ctx context.Context, peripheralId string) error {
	db := database.Get(initTable)

	_, err := db.ExecContext(ctx, `DELETE FROM PeripheralEvents WHERE PeripheralId = ?;`, peripheralId)

	return err
}

func MostRecentDeploymentEvents(ctx context.Context, deploymentId string) ([]api.MeasurementEvent, error) {
	db := database.Get(initTable)

	events, err := ScanEvents(db.QueryContext(
		ctx,
		`SELECT * from PeripheralEvents WHERE (PeripheralId, Timestamp)
		IN (
			SELECT PeripheralId, MAX(Timestamp) AS Timestamp
			FROM PeripheralEvents
			WHERE DeploymentId = ?
			GROUP BY PeripheralId
		);	`,
		&deploymentId,
	))

	if err != nil {
		log.Printf("Failed to get events out of db, error %v", err)
	}

	log.Printf("Found %v events", len(events))

	return events, err
}

func initTable(ctx context.Context, pool *sql.DB) error {
	_, peripheralEventsTableCreateErr := pool.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS PeripheralEvents (
        Id varchar(36) PRIMARY KEY NOT NULL,
        PeripheralId varchar(36) NOT NULL,
        DeploymentId varchar(36) NOT NULL,
        Value float NOT NULL,
        Timestamp timestamp NOT NULL
        );`,
	)
	if peripheralEventsTableCreateErr != nil {
		log.Printf(
			"Failed to create peripheral events table. error: %v",
			peripheralEventsTableCreateErr,
		)
	}
	return peripheralEventsTableCreateErr
}

func ConvertPBTimeToTime(pbtime *timestamppb.Timestamp) (time.Time, error) {
	time, err := ptypes.Timestamp(pbtime)
	if err != nil {
		log.Printf("Failed to convert timestamp to sql consumable type, error %v", err)
	}

	return time, err
}
