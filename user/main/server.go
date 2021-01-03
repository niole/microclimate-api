package main

import (
	api "api/protobuf"
	"api/shared/server"
	"api/user/service"
	"google.golang.org/grpc"
)

func run(grpcServer *grpc.Server) {
	api.RegisterUserServiceServer(
		grpcServer,
		new(service.UserServiceServer),
	)
}

func main() {
	server.BaseServer(run)
}
