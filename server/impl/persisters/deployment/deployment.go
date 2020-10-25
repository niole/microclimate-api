package deployment

import (
	api "api/protobuf"
	connector "api/server/impl/database"
	"context"
	sql "database/sql"
	"log"
)

var statusMap = map[string]api.Deployment_Status{
	"unreachable":   api.Deployment_UNREACHABLE,
	"starting_up":   api.Deployment_STARTING_UP,
	"failed":        api.Deployment_FAILED,
	"running":       api.Deployment_RUNNING,
	"shutting_down": api.Deployment_SHUTTING_DOWN,
}

func ScanOneDeployment(result *sql.Row, saveToDeployment *api.Deployment) error {
	var id string
	var ownerUserId string
	var domain string
	var status string
	err := result.Scan(&id, &ownerUserId, &domain, &status)

	*saveToDeployment = api.Deployment{
		Id:          id,
		OwnerUserId: ownerUserId,
		Domain:      domain,
		Status:      statusMap[status],
	}

	return err
}

// creates deployment for user
func CreateDeployment(ctx context.Context, newDeployment *api.NewDeployment) (*api.Deployment, error) {
	db := connector.GetConnectionPool()

	_, err := db.ExecContext(ctx, `INSERT INTO Deployments (Id, OwnerUserId, Domain, Status) VALUES (UUID(), ?, ?, ?)`,
		&newDeployment.OwnerUserId,
		"localhost",
		"unreachable",
	)

	if err != nil {
		log.Printf(
			"Failed to insert new deployment: %v, err: %v",
			&newDeployment,
			err,
		)
		return nil, err
	}

	var deployment api.Deployment

	err = ScanOneDeployment(db.QueryRowContext(
		ctx,
		"SELECT * FROM Deployments WHERE OwnerUserId = ? LIMIT 1;",
		&newDeployment.OwnerUserId,
	), &deployment)

	return &deployment, err
}
