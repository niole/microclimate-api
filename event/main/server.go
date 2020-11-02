package main

import (
	"api/event/service"
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
	flag.Parse()
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterPeripheralMeasurementEventServiceServer(
		grpcServer,
		new(service.PeripheralEventService),
	)
	grpcServer.Serve(lis)
}
