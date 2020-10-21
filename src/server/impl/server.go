package main

import (
	api "api/protobuf"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port int = 6000

type deploymentServiceServer struct {
	api.UnimplementedDeploymentManagementServiceServer
}

func (s deploymentServiceServer) CreateDeployment(ctx context.Context, newDeployment *api.NewDeployment) (*api.Deployment, error) {
	deployment := &api.Deployment{
		Id:          1,
		OwnerUserId: newDeployment.OwnerUserId,
		Domain:      "localhost",
		Status:      api.Deployment_UNREACHABLE,
	}

	return deployment, nil
}

func (s deploymentServiceServer) GetDeployment(ctx context.Context, in *api.GetDeploymentRequest) (*api.Deployment, error) {
	deployment := &api.Deployment{
		Id:          1,
		OwnerUserId: in.OwnerUserId,
		Domain:      "localhost",
		Status:      api.Deployment_UNREACHABLE,
	}
	return deployment, nil
}

func (s deploymentServiceServer) RemoveDeployment(ctx context.Context, in *api.RemoveDeploymentRequest) (*api.Empty, error) {
	return &api.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDeploymentManagementServiceServer(
		grpcServer,
		deploymentServiceServer{},
	)
	grpcServer.Serve(lis)
}
