package main

import (
	"api/event/service"
	api "api/protobuf"
	"api/shared/server"
	"google.golang.org/grpc"
)

func run(grpcServer *grpc.Server) {
	api.RegisterPeripheralMeasurementEventServiceServer(
		grpcServer,
		new(service.PeripheralEventService),
	)
}

func main() {
	server.BaseServer(run)
}
