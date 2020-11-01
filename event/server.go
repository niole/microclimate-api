package main

import (
	"api/event"
	api "api/protobuf"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port int

func init() {
	flag.IntVar(&port, "serverPort", 6002, "Port for eventservice server")
}

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
