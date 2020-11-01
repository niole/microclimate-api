package main

import (
	api "api/protobuf"
	user "api/server/impl/user"
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
	api.RegisterUserServiceServer(
		grpcServer,
		new(user.UserServiceServer),
	)
	grpcServer.Serve(lis)
}
