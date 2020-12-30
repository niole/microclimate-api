package main

import (
	"api/clients"
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
var host string

func init() {
	flag.IntVar(&port, "serverPort", 8080, "Service's port")
	flag.StringVar(&host, "host", "localhost", "Service's host")
	flag.Parse()
}

func newUuidString() string {
	return uuid.New().String()
}

func testEvent(conn grpc.ClientConnInterface) {
	client := api.NewPeripheralMeasurementEventServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	peripheralId := "p1"
	deploymentId := "deploymentuid"

	for i := 0; i < 5; i++ {
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

	start, err := ptypes.TimestampProto(time.Date(2019, 12, 31, 19, 0, 0, 0, time.UTC))
	end, err := ptypes.TimestampProto(time.Now())
	log.Print(start)
	log.Print(end)
	if err != nil {
		log.Fatalf("The times are bad, err %v", err)
	}
	eventStreamClient, streamError := client.FilterEvents(ctx, &api.MeasurementEventFilterRequest{
		PeripheralId: peripheralId,
		DeploymentId: deploymentId,
		StartTime:    start,
		EndTime:      end,
	})

	if streamError != nil {
		log.Fatalf("Failed to start receiving events, error %v", streamError)
	}

	log.Print("HEY")

	for {
		newEvent, err := eventStreamClient.Recv()
		log.Print(err)
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	log.Print("starting deployment test")

	ownerUserId := "nioleuid"

	log.Print("Creating deployment")
	d, err := client.CreateDeployment(ctx, &api.NewDeployment{
		OwnerUserId: ownerUserId,
		Name:        "gmy apartment 123",
	})
	log.Print("AFTER Creating deployment")

	////_, err = client.CreateDeployment(ctx, &api.NewDeployment{
	////	OwnerUserId: ownerUserId,
	////})

	log.Print(d)
	log.Print(err)

	log.Print("Get user deployments")

	deploymentStreamClient, streamError := client.GetDeploymentsForUser(ctx, &api.GetDeploymentsForUserRequest{
		UserId: ownerUserId,
	})

	if streamError != nil {
		log.Fatalf("Failed to start receiving deployments, error %v", streamError)
	}

	for {
		newEvent, err := deploymentStreamClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive single deployment, error %v", err)
			break
		}
		if newEvent != nil {
			log.Print(newEvent)
		} else {
			log.Print("DONE")
			break
		}
	}

	gottenD, err := client.GetDeployment(ctx, &api.GetDeploymentRequest{
		OwnerUserId:  ownerUserId,
		DeploymentId: "deploymentuid",
	})

	log.Print(gottenD)
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

	email := "newuser@gmail.com"
	u, err := client.CreateUser(ctx, &api.NewUser{Email: email, Name: "newuserName"})
	log.Print(u)
	log.Print(err)

	u, err = client.GetUserByEmail(ctx, &api.GetUserByEmailRequest{Email: email})
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

	p, err := client.LinkHardware(ctx, &api.LinkHardwareRequest{HardwareId: newUuidString(), PeripheralId: "fakep2"})

	log.Print(p)
	log.Print(err)

	p, err = client.LinkHardware(ctx, &api.LinkHardwareRequest{HardwareId: newUuidString(), PeripheralId: "p2"})

	log.Print(p)
	log.Print(err)

	ownerUserId := "nioleuid"
	deploymentId := "deploymentuid"
	p, err = client.CreatePeripheral(ctx, &api.NewPeripheral{
		OwnerUserId:  ownerUserId,
		DeploymentId: deploymentId,
		HardwareId: &api.NullableString{
			Kind: &api.NullableString_Null{},
		},
		Type: api.NewPeripheral_PARTICLE,
		Unit: "PM2.5",
		Name: "with null id",
	})
	newP := p

	//client.CreatePeripheral(ctx, &api.NewPeripheral{
	//	OwnerUserId:  ownerUserId,
	//	DeploymentId: deploymentId,
	//	HardwareId: &api.NullableString{
	//		Kind: &api.NullableString_Data{
	//			Data: newUuidString(),
	//		},
	//	},
	//	Type: api.NewPeripheral_PARTICLE,
	//	Unit: "PM2.5",
	//	Name: "livingroom",
	//})

	log.Print(p)
	log.Print(err)

	p, err = client.GetPeripheral(ctx, &api.GetPeripheralRequest{
		PeripheralId: newP.Id,
	})

	log.Print(p)
	log.Print(err)

	p, err = client.EditPeripheral(ctx, &api.EditPeripheralRequest{
		NewName: &api.NullableString{
			Kind: &api.NullableString_Data{
				Data: "NAMENAME",
			},
		},
		NewType: &api.NullablePType{
			Kind: &api.NullablePType_Data{
				Data: api.Peripheral_THERMAL,
			},
		},
		PeripheralId: newP.Id,
	})

	log.Print(p)
	log.Print(err)
	p, err = client.GetPeripheral(ctx, &api.GetPeripheralRequest{
		PeripheralId: newP.Id,
	})
	log.Print(p)
	log.Print(err)

	_, err = client.RemovePeripheral(ctx, newP)

	log.Print(err)
}

func main() {
	log.Println(quote.Go())
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	log.Println(serverAddr)
	conn, err := clients.ClientConnection(serverAddr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	testDeployment(conn)

	//testEvent(conn)
	//testPeripheral(conn)
	//testUser(conn)
}
