package clients

import (
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"os"
)

var (
	serviceAccountJson   = os.Getenv("SERVICE_ACCOUNT_JSON")
	eventServerAddr      = os.Getenv("EVENT_SERVER_ADDR")
	peripheralServerAddr = os.Getenv("PERIPHERAL_SERVER_ADDR")
)

type myPerRPC struct {
	credentials.PerRPCCredentials
}

func (m myPerRPC) RequireTransportSecurity() bool { return false }

func JwtOnlyClientConnection(serverAddr string) (*grpc.ClientConn, error) {
	scope := "read-write"
	perRPC, err := oauth.NewServiceAccountFromFile(serviceAccountJson, scope)

	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(myPerRPC{perRPC}),
	)

	if err != nil {
		log.Printf("fail to dial: %v", err)
		return nil, err
	}
	return conn, err
}

func InsecureClientConnection(serverAddr string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Printf("fail to dial: %v", err)
		return nil, err
	}
	return conn, err
}

func ClientConnection(serverAddr string) (*grpc.ClientConn, error) {
	pool, err := x509.SystemCertPool()
	creds := credentials.NewClientTLSFromCert(pool, "")
	scope := "read-write"

	perRPC, err := oauth.NewServiceAccountFromFile(serviceAccountJson, scope)

	if err != nil {
		log.Printf("Failed to create jwt %v", err)
	}
	fmt.Println(perRPC)
	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(perRPC),
	)

	if err != nil {
		log.Printf("fail to dial: %v", err)
		return nil, err
	}
	return conn, err
}

func PeripheralClientConnection() (*grpc.ClientConn, error) {
	addr := peripheralServerAddr
	if peripheralServerAddr == "" {
		addr = "0.0.0.0:6001"
	}
	return ClientConnection(addr)
}

func EventsClientConnection() (*grpc.ClientConn, error) {
	addr := eventServerAddr
	if eventServerAddr == "" {
		addr = "0.0.0.0:6004"
	}
	return ClientConnection(addr)
}
