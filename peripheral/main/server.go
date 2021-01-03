package main

import (
	"api/peripheral/service"
	api "api/protobuf"
	"api/shared/server"
	"google.golang.org/grpc"
)

func run(grpcServer *grpc.Server) {
	api.RegisterPeripheralManagementServiceServer(
		grpcServer,
		new(service.PeripheralManagementService),
	)
}

func main() {
	server.BaseServer(run)
}
