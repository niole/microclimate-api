package main

import (
	api "api/protobuf"
	//deployment "api/server/impl/deployment"
	peripherals "api/server/impl/peripheral"
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
	// TODO what should we do about the multiple servers situation
	//api.RegisterDeploymentManagementServiceServer(
	//	grpcServer,
	//	new(deployment.DeploymentManagementService),
	//)
	api.RegisterPeripheralManagementServiceServer(
		grpcServer,
		new(peripherals.PeripheralManagementService),
	)
	grpcServer.Serve(lis)
}
