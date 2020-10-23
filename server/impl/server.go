package main

import (
	api "api/protobuf"
	deployment "api/server/impl/deployment"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port int = 6000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDeploymentManagementServiceServer(
		grpcServer,
		new(deployment.DeploymentManagementService),
	)
	grpcServer.Serve(lis)
}
