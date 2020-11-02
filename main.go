package main

import (
	api "api/protobuf"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
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

func testEvent(conn grpc.ClientConnInterface) {
	client := api.NewPeripheralMeasurementEventServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	peripheralId := newUuidString()
	deploymentId := newUuidString()

	for i := 0; i < 5; i++ {
		thetime, err := ptypes.TimestampProto(
			time.Date(2020, 10, 31, 19, 2*i, 0, 0, time.UTC),
		)
		_, err = client.SendEvent(
			ctx,
			&api.NewMeasurementEvent{
				PeripheralId: peripheralId,
				DeploymentId: deploymentId,
				Value:        80,
				TimeStamp:    thetime,
			},
		)

		log.Printf("Sent a request did it fail %v", err)
	}
	start, err := ptypes.TimestampProto(time.Date(2020, 10, 31, 19, 0, 0, 0, time.UTC))
	end, err := ptypes.TimestampProto(time.Date(2020, 10, 31, 19, 7, 0, 0, time.UTC))
	if err != nil {
		log.Fatalf("The times are bad, err %v", err)
	}
	eventStreamClient, streamError := client.FilterEvents(ctx, &api.MeasurementEventFilterRequest{
		PeripheralIds: []string{peripheralId},
		DeploymentId:  deploymentId,
		StartTime:     start,
		EndTime:       end,
	})

	if streamError != nil {
		log.Fatalf("Failed to start receiving events, error %v", streamError)
	}

	for {
		newEvent, err := eventStreamClient.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive single event, error %v", err)
			return
		}
		if newEvent != nil {
			log.Print(newEvent)
		} else {
			log.Print("DONE")
			return
		}
	}
}

func testDeployment(conn grpc.ClientConnInterface) {
	client := api.NewDeploymentManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Print("starting deployment test")
	d, err := client.CreateDeployment(ctx, &api.NewDeployment{
		OwnerUserId: newUuidString(),
	})

	log.Print(d)
	log.Print(err)

	d, err = client.GetDeployment(ctx, &api.GetDeploymentRequest{
		OwnerUserId:  d.OwnerUserId,
		DeploymentId: d.Id,
	})

	log.Print(d)
	log.Print(err)

	_, err = client.RemoveDeployment(ctx, &api.RemoveDeploymentRequest{
		DeploymentId: d.Id,
	})

	log.Print(err)
}

func testPeripheral(conn grpc.ClientConnInterface) {
	client := api.NewPeripheralManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Print("starting peripheral test")

	p, err := client.CreatePeripheral(ctx, &api.NewPeripheral{
		OwnerUserId:  newUuidString(),
		DeploymentId: newUuidString(),
		HardwareId:   newUuidString(),
		Type:         api.NewPeripheral_PARTICLE,
	})

	log.Print(p)
	log.Print(err)

	peripheralId := p.Id

	p, err = client.GetPeripheral(ctx, &api.GetPeripheralRequest{
		PeripheralId: peripheralId,
	})

	log.Print(p)
	log.Print(err)

	_, err = client.RemovePeripheral(ctx, p)

	log.Print(err)
}

func main() {
	log.Println(quote.Go())
	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	serverAddr := fmt.Sprintf("localhost:%d", 6003)
	log.Println(serverAddr)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	//testEvent(conn)
	testPeripheral(conn)

	//deploymentId := "c1e2052c-06f3-4a2b-a8a5-73ac7651c022"
	//peripheralId := "51e6f06e-1a68-11eb-888c-0242ac110002"

	//found, err := client.GetPeripheral(ctx, &api.GetPeripheralRequest{
	//	PeripheralId: peripheralId,
	//})

	//log.Print(found)
	//log.Print(err)
}
