package server

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type serverFlags struct {
	enableReflection bool
	port             int
}

func parseFlags() serverFlags {
	enableReflection := flag.Bool("withReflection", false, "Whether or not to enable reflection")
	port := flag.Int("serverPort", 8080, "Port for server")
	flag.Parse()

	return serverFlags{*enableReflection, *port}
}

func BaseServer(registrar func(*grpc.Server)) {
	flags := parseFlags()

	address := fmt.Sprintf("0.0.0.0:%d", flags.port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Listening at %v", address)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	if flags.enableReflection {
		reflection.Register(grpcServer)
		log.Print("Reflection enabled")
	}

	registrar(grpcServer)

	grpcServer.Serve(lis)
}
