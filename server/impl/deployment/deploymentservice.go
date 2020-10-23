package deployment

import (
	api "api/protobuf"
	"context"
)

type DeploymentManagementService struct {
	api.UnimplementedDeploymentManagementServiceServer
}

func (s DeploymentManagementService) CreateDeployment(ctx context.Context, newDeployment *api.NewDeployment) (*api.Deployment, error) {
	deployment := &api.Deployment{
		Id:          1,
		OwnerUserId: newDeployment.OwnerUserId,
		Domain:      "localhost",
		Status:      api.Deployment_UNREACHABLE,
	}

	return deployment, nil
}

func (s DeploymentManagementService) GetDeployment(ctx context.Context, in *api.GetDeploymentRequest) (*api.Deployment, error) {
	deployment := &api.Deployment{
		Id:          1,
		OwnerUserId: in.OwnerUserId,
		Domain:      "localhost",
		Status:      api.Deployment_UNREACHABLE,
	}
	return deployment, nil
}

func (s DeploymentManagementService) RemoveDeployment(ctx context.Context, in *api.RemoveDeploymentRequest) (*api.Empty, error) {
	return &api.Empty{}, nil
}
