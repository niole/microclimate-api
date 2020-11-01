package persister

import (
	api "api/protobuf"
	"api/server/impl/database"
	"context"
	"log"
)

// creates deployment for user
func CreateDeployment(ctx context.Context, newDeployment *api.NewDeployment) (*api.Deployment, error) {
	db := database.Get(init)

	_, err := db.ExecContext(
		ctx,
		`INSERT INTO Deployments (Id, OwnerUserId, Domain, Status) VALUES (UUID(), ?, ?, ?)`,
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

func GetDeployment(ctx context.Context, in *api.GetDeploymentRequest) (*api.Deployment, error) {
	db := database.Get(ctx, init)
	var deployment api.Deployment

	err := ScanOneDeployment(db.QueryRowContext(
		ctx,
		"SELECT * FROM Deployments WHERE OwnerUserId = ? AND Id = ? LIMIT 1;",
		&in.OwnerUserId,
		&in.DeploymentId,
	), &deployment)

	return &deployment, err
}

func RemoveDeployment(ctx context.Context, in *api.RemoveDeploymentRequest) error {
	db := database.Get(init)

	_, err := db.ExecContext(ctx, `DELETE FROM Deployments WHERE ID = ? `, &in.DeploymentId)

	return err
}

func init(ctx context.Context, pool *sql.DB) error {
	_, deploymentTableCreateErr := pool.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS Deployments (
                Id varchar(36) PRIMARY KEY NOT NULL,
                OwnerUserId varchar(36) NOT NULL,
                Domain varchar(255) NOT NULL,
                Status ENUM('unreachable', 'starting_up', 'failed', 'running', 'shutting_down') NOT NULL
            );`,
	)
	if deploymentTableCreateErr != nil {
		log.Fatalf("Failed to create deployments table. error: %v", deploymentTableCreateErr)
	}
	return deploymentTableCreateErr
}
