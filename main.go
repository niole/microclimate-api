package main

import (
	api "api/protobuf"
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc"
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

	//perfIds := []string{"1e8b843e-2a31-11eb-96aa-0242ac1d0002", "1e8d604c-2a31-11eb-96aa-0242ac1d0002", "6c376575-2b49-11eb-96aa-0242ac1d0002"}
	peripheralId := "1cee9bb1-aa12-4a04-80c1-93c4e27bfd52"
	deploymentId := "45e1ecaa-3324-11eb-a9e5-0242c0a80002"

	for i := 0; i < 500; i++ {
		thetime, err := ptypes.TimestampProto(
			time.Date(2020, 10, 31, 19, 2*i, 0, 0, time.UTC),
		)
		_, err = client.SendEvent(
			ctx,
			&api.NewMeasurementEvent{
				PeripheralId: peripheralId,
				DeploymentId: deploymentId,
				Value:        0,
				TimeStamp:    thetime,
			},
		)
		if err != nil {
			log.Printf("Sent the requests did they fail %v", err)
		}
	}

	//start, err := ptypes.TimestampProto(time.Date(1979, 10, 31, 19, 0, 0, 0, time.UTC))
	//end, err := ptypes.TimestampProto(time.Now())
	//log.Print(start)
	//log.Print(end)
	//if err != nil {
	//	log.Fatalf("The times are bad, err %v", err)
	//}
	//eventStreamClient, streamError := client.FilterEvents(ctx, &api.MeasurementEventFilterRequest{
	//	PeripheralId: perfIds,
	//	DeploymentId: deploymentId,
	//	StartTime:    start,
	//	EndTime:      end,
	//})

	//if streamError != nil {
	//	log.Fatalf("Failed to start receiving events, error %v", streamError)
	//}

	//log.Print("HEY")

	//for {
	//	newEvent, err := eventStreamClient.Recv()
	//	log.Print(err)
	//	if err == io.EOF {
	//		return
	//	}
	//	if err != nil {
	//		log.Fatalf("Failed to receive single event, error %v", err)
	//		return
	//	}
	//	if newEvent != nil {
	//		log.Print(newEvent)
	//	} else {
	//		log.Print("DONE")
	//		return
	//	}
	//}
}

func testDeployment(conn grpc.ClientConnInterface) {
	client := api.NewDeploymentManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Print("starting deployment test")

	ownerUserId := "2a1fb3e7-3324-11eb-a9e5-0242c0a80002"

	d, err := client.CreateDeployment(ctx, &api.NewDeployment{
		OwnerUserId: ownerUserId,
		Name:        "my apartment 123",
	})

	//_, err = client.CreateDeployment(ctx, &api.NewDeployment{
	//	OwnerUserId: ownerUserId,
	//})

	log.Print(d)
	log.Print(err)

	//deploymentStreamClient, streamError := client.GetDeploymentsForUser(ctx, &api.GetDeploymentsForUserRequest{
	//	UserId: ownerUserId,
	//})

	//if streamError != nil {
	//	log.Fatalf("Failed to start receiving deployments, error %v", streamError)
	//}

	//for {
	//	newEvent, err := deploymentStreamClient.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("Failed to receive single deployment, error %v", err)
	//		break
	//	}
	//	if newEvent != nil {
	//		log.Print(newEvent)
	//	} else {
	//		log.Print("DONE")
	//		break
	//	}
	//}

	//d, err = client.GetDeployment(ctx, &api.GetDeploymentRequest{
	//	OwnerUserId:  d.OwnerUserId,
	//	DeploymentId: d.Id,
	//})

	//log.Print(d)
	//log.Print(err)

	//_, err = client.RemoveDeployment(ctx, &api.RemoveDeploymentRequest{
	//	DeploymentId: d.Id,
	//})

	//log.Print(err)
}

func testUser(conn grpc.ClientConnInterface) {
	client := api.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Print("starting user test")

	u, err := client.CreateUser(ctx, &api.NewUser{Email: "niolenelson@gmail.com", Name: "niole"})
	log.Print(u)
	log.Print(err)

	//u, err := client.GetUserByEmail(ctx, &api.GetUserByEmailRequest{Email: "niolenelson@gmail"})
	//log.Print(u)
	//log.Print(err)

	//u, err = client.UpdateUserEmail(ctx, &api.UpdateUserEmailRequest{UserId: u.Id, Email: "newemail"})
	//log.Print(u)
	//log.Print(err)

	//_, err = client.RemoveUser(ctx, &api.RemoveUserRequest{UserId: u.Id})
	//log.Print(err)

}

func testPeripheral(conn grpc.ClientConnInterface) {
	client := api.NewPeripheralManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Print("starting peripheral test")

	ownerUserId := "2a1fb3e7-3324-11eb-a9e5-0242c0a80002"
	deploymentId := "45e1ecaa-3324-11eb-a9e5-0242c0a80002"

	p, err := client.CreatePeripheral(ctx, &api.NewPeripheral{
		OwnerUserId:  ownerUserId,
		DeploymentId: deploymentId,
		HardwareId:   newUuidString(),
		Type:         api.NewPeripheral_PARTICLE,
		Unit:         "PM2.5",
		Name:         "garage particle sensor",
	})

	client.CreatePeripheral(ctx, &api.NewPeripheral{
		OwnerUserId:  ownerUserId,
		DeploymentId: deploymentId,
		HardwareId:   newUuidString(),
		Type:         api.NewPeripheral_PARTICLE,
		Unit:         "PM2.5",
		Name:         "livingroom",
	})

	log.Print(p)
	log.Print(err)

	//peripheralId := p.Id

	//p, err = client.GetPeripheral(ctx, &api.GetPeripheralRequest{
	//	PeripheralId: peripheralId,
	//})

	//log.Print(p)
	//log.Print(err)

	//_, err = client.RemovePeripheral(ctx, p)

	//log.Print(err)
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

	//testDeployment(conn)

	testEvent(conn)
	//testPeripheral(conn)
	//testUser(conn)
}
