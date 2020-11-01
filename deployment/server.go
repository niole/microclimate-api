package main

import (
	"api/deployment"
	api "api/protobuf"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port int

func init() {
	flag.IntVar(&port, "serverPort", 6001, "Port for deploymentservice server")
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
		new(deployment.DeploymentManagementService),
	)
	grpcServer.Serve(lis)
}
