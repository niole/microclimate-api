package main

import (
	api "api/protobuf"
	"api/server/impl/event"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// TODO get from env
var port int = 6002

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterPeripheralMeasurementEventServiceServer(
		grpcServer,
		new(event.PeripheralEventService),
	)
	grpcServer.Serve(lis)
}
