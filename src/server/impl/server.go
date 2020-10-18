package main

import (
	pb "api/protobuf/tutorial/tutorialpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port int = 10000

type searchServiceServer struct {
}

func (s *searchServiceServer) Search(context.Context, *pb.AddressBook) (*pb.Person, error) {
	return &pb.Person{
		Name:  "niole",
		Id:    123,
		Email: "e@mail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123", Type: pb.Person_MOBILE},
		},
	}, nil
}

func (s *searchServiceServer) mustEmbedUnimplementedSearchServiceServer() {
}

func newServer() searchServiceServer {
	return searchServiceServer{}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", &port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
