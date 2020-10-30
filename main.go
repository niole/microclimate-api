package main

import (
	api "api/protobuf"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
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

	deploymentId := newUuidString()
	ownerUserId := newUuidString()

	_, pErr := client.CreatePeripheral(ctx, &api.NewPeripheral{
		OwnerUserId:  ownerUserId,
		DeploymentId: deploymentId,
		HardwareId:   "duniqueindeployemnthwid",
		Type:         api.NewPeripheral_THERMAL,
	})

	_, pErr = client.CreatePeripheral(ctx, &api.NewPeripheral{
		OwnerUserId:  ownerUserId,
		DeploymentId: deploymentId,
		HardwareId:   "uniqueindeployemnthwid",
		Type:         api.NewPeripheral_THERMAL,
	})

	if pErr != nil {
		log.Printf("add peripheral err %v", pErr)
	}

	//_, removeErr := client.RemovePeripheral(ctx, p)

	//if removeErr != nil {
	//	log.Printf("remove ed one, errerr %v", removeErr)
	//}

	peripheralsClient, streamError := client.GetDeploymentPeripherals(
		ctx,
		&api.GetDeploymentPeripheralsRequest{
			DeploymentId: deploymentId,
		},
	)

	log.Print("Receiveing stream peripherals")

	if streamError != nil {
		log.Fatalf("Failed to start receiving peripherals, error %v", streamError)
	}

	for {
		newPeripheral, receiveError := peripheralsClient.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive single peripheral, error %v", receiveError)
			return
		}
		if newPeripheral != nil {
			log.Print(newPeripheral)
		} else {
			log.Print("DONE")
			return
		}
	}

}
