package event

import (
	api "api/protobuf"
	connection "api/server/impl/database"
	"context"
	"github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	"log"
	"strings"
	"time"
)

func ConvertPBTimeToTime(pbtime *timestamppb.Timestamp) (time.Time, error) {
	time, err := ptypes.Timestamp(pbtime)
	if err != nil {
		log.Fatalf("Failed to convert timestamp to sql consumable type, error %v", err)
	}

	return time, err
}

func SaveEvent(ctx context.Context, event *api.NewMeasurementEvent) error {
	db := connection.GetConnectionPool()

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
	db := connection.GetConnectionPool()

	in := strings.Join(request.PeripheralIds, ",")

	endtime, _ := ConvertPBTimeToTime(request.EndTime)
	starttime, _ := ConvertPBTimeToTime(request.StartTime)

	log.Print(in)

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
