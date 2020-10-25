package deployment

import (
	api "api/protobuf"
	persister "api/server/impl/persisters/deployment"
	"context"
	"log"
	"time"
)

type DeploymentManagementService struct {
	api.UnimplementedDeploymentManagementServiceServer
}

func (s DeploymentManagementService) CreateDeployment(ctx context.Context, newDeployment *api.NewDeployment) (*api.Deployment, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	deployment, err := persister.CreateDeployment(cancellableCtx, newDeployment)

	if err != nil {
		log.Printf(
			"Failed to get deployment out of db. requested deployment: %v, err: %v",
			&newDeployment,
			err,
		)
	}

	return deployment, err
}

func (s DeploymentManagementService) GetDeployment(ctx context.Context, in *api.GetDeploymentRequest) (*api.Deployment, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	deployment, err := persister.GetDeployment(cancellableCtx, in)

	if err != nil {
		log.Printf(
			"Failed to get deployment out of db. requested user and deployment: %v, err: %v",
			&in,
			err,
		)
	}

	return deployment, err
}

func (s DeploymentManagementService) RemoveDeployment(ctx context.Context, in *api.RemoveDeploymentRequest) (*api.Empty, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := persister.RemoveDeployment(cancellableCtx, in)

	if err != nil {
		log.Printf("Failed to remove deployment %v", in)
	}

	return &api.Empty{}, err
}
