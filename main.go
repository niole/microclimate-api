package main

import (
	api "api/protobuf"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
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
	client := api.NewPeripheralMeasurementEventServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = client.SendEvent(
		ctx,
		&api.MeasurementEvent{
			PeripheralId: "51e6f06e-1a68-11eb-888c-0242ac110002",
			DeploymentId: "c1e2052c-06f3-4a2b-a8a5-73ac7651c022",
			Value:        80,
			TimeStamp:    ptypes.TimestampNow(),
		},
	)

	log.Printf("Sent a request did it fail %v", err)

	//if streamError != nil {
	//	log.Fatalf("Failed to start receiving peripherals, error %v", streamError)
	//}

	//for {
	//	newPeripheral, receiveError := peripheralsClient.Recv()
	//	if err == io.EOF {
	//		return
	//	}
	//	if err != nil {
	//		log.Fatalf("Failed to receive single peripheral, error %v", receiveError)
	//		return
	//	}
	//	if newPeripheral != nil {
	//		log.Print(newPeripheral)
	//	} else {
	//		log.Print("DONE")
	//		return
	//	}
	//}

}
