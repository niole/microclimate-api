package clients

import (
	"google.golang.org/grpc"
	"log"
)

var (
	peripheralServerAddr = os.Getenv("PERIPHERAL_SERVER_ADDR")
)

func clientConnection(serverAddr string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Printf("fail to dial: %v", err)
		return nil, err
	}
	return conn, err
}

func PeripheralClientConnection() (*grpc.ClientConn, error) {
	addr := peripheralServerAddr
	if peripheralServerAddr == nil {
		addr = "0.0.0.0:6001"
	}
	return clientConnection(addr)
}
