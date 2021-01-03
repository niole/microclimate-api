package main

import (
	"api/deployment/service"
	api "api/protobuf"
	"api/shared/server"
	"google.golang.org/grpc"
)

func run(grpcServer *grpc.Server) {
	api.RegisterDeploymentManagementServiceServer(
		grpcServer,
		new(service.DeploymentManagementService),
	)
}

func main() {
	server.BaseServer(run)
}
