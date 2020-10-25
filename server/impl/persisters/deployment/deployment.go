package deployment

import (
	api "api/protobuf"
	connector "api/server/impl/database"
	"context"
	"log"
)

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
