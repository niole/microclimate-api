package main

import (
	api "api/protobuf"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"rsc.io/quote"
	"time"
)

func newUuidString() string {
	return uuid.New().String()
}

func main() {
	log.Println(quote.Go())
	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	serverAddr := fmt.Sprintf("localhost:%d", 6000)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewPeripheralManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	p, pErr := client.CreatePeripheral(ctx, &api.NewPeripheral{
		OwnerUserId:  newUuidString(),
		DeploymentId: newUuidString(),
		HardwareId:   "uniqueindeployemnthwid",
		Type:         api.NewPeripheral_THERMAL,
	})

	log.Printf("%v, err %v", p, pErr)

	_, removeErr := client.RemovePeripheral(ctx, p)

	log.Printf("remove err %v", removeErr)
}
