package persister

import (
	api "api/protobuf"
	"api/shared/database"
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	"log"
	"math"
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

	// TODO the averaging doesn't average right
	// the average looks like its a kind of rolling average, when it should be a smooth average
	// also the last window looks like it doesn't get averaged
	// also the UI should treat "end" dates as "the end of the whole day"
	// also doing this operation over thousands of events takes a while
	// 5 second deadline exceeds (easily) when 15000 events are
	// processed by the filter query
	endtime, _ := ConvertPBTimeToTime(request.EndTime)
	starttime, _ := ConvertPBTimeToTime(request.StartTime)

	log.Printf(
		"Getting events. start time %v, end time %v, peripheral id %v, deploymentid %v",
		starttime,
		endtime,
		request.PeripheralId,
		request.DeploymentId,
	)

	var count float64
	countRows, err := db.QueryContext(
		ctx,
		`SELECT COUNT(*) as count FROM PeripheralEvents WHERE
		Timestamp BETWEEN ? AND ?
		AND DeploymentId = ?
		AND PeripheralId = ?
		;`,
		starttime,
		endtime,
		&request.DeploymentId,
		&request.PeripheralId,
	)
	for countRows.Next() {
		err = countRows.Scan(&count)
	}

	if err != nil {
		log.Printf("Failed to count events, %v", err)
	} else {
		maxBuckets := 100.0
		windowSize := int(math.Max(math.Ceil(count/maxBuckets), 1.0))
		log.Printf("window size %v", windowSize)
		events, err := ScanAggregatedEvents(db.QueryContext(
			ctx,
			`SELECT * FROM (
				SELECT *,
				AVG(Value) OVER (ORDER BY Timestamp ROWS ? PRECEDING) AS average,
				ROW_NUMBER() OVER (ORDER BY Timestamp) AS n FROM PeripheralEvents
			) x WHERE n % ? = 0
			AND Timestamp BETWEEN ? AND ?
			AND DeploymentId = ?
			AND PeripheralId = ?
			;`,
			windowSize,
			windowSize,
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

	return make([]api.MeasurementEvent, 0), err
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
		);  `,
		&deploymentId,
	))

	if err != nil {
		log.Printf("Failed to get events out of db, error %v", err)
	}

	log.Printf("Found %v events", len(events))

	return events, err
}

func initTable(ctx context.Context, pool *sql.DB) error {
	// make sure floats are formatted to two decimal points
	// get rid of deployment id
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
