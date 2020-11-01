package main

import (
	api "api/protobuf"
	"api/user"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port int

func init() {
	flag.IntVar(&port, "serverPort", 6004, "Port for userservice server")
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterUserServiceServer(
		grpcServer,
		new(peripheral.UserServiceServer),
	)
	grpcServer.Serve(lis)
}
