package main

import (
	api "api/protobuf"
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
	"log"
	"rsc.io/quote"
	"time"
)

var port int

func init() {
	flag.IntVar(&port, "serverPort", 8080, "Port client")
	flag.Parse()
}

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

func testUser(conn grpc.ClientConnInterface) {
	client := api.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Print("starting user test")

	u, err := client.CreateUser(ctx, &api.NewUser{Email: "oldemail"})
	log.Print(u)
	log.Print(err)

	u, err = client.GetUserByEmail(ctx, &api.GetUserByEmailRequest{Email: u.Email})
	log.Print(u)
	log.Print(err)

	u, err = client.UpdateUserEmail(ctx, &api.UpdateUserEmailRequest{UserId: u.Id, Email: "newemail"})
	log.Print(u)
	log.Print(err)

	_, err = client.RemoveUser(ctx, &api.RemoveUserRequest{UserId: u.Id})
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
	serverAddr := fmt.Sprintf("localhost:%d", port)
	log.Println(serverAddr)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	//testEvent(conn)
	//testPeripheral(conn)
	testUser(conn)
}
