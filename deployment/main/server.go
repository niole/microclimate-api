package main

import (
	"api/deployment/service"
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
	log.Printf("Starting deployment serve at port %v", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDeploymentManagementServiceServer(
		grpcServer,
		new(service.DeploymentManagementService),
	)
	grpcServer.Serve(lis)
}
