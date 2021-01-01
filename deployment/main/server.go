package main

import (
	"api/deployment/service"
	api "api/protobuf"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var port int
var enableReflection bool

func init() {
	flag.BoolVar(&enableReflection, "withReflection", false, "Whether or not to enable reflection")
	flag.IntVar(&port, "serverPort", 6001, "Port for deploymentservice server")
	flag.Parse()
}

func main() {
	log.Printf("Starting deployment serve at port %v", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	if enableReflection {
		reflection.Register(grpcServer)
	}

	api.RegisterDeploymentManagementServiceServer(
		grpcServer,
		new(service.DeploymentManagementService),
	)
	grpcServer.Serve(lis)
}
