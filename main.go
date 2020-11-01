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
	client := api.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//startemail := "catdog"
	newemail := "dog"
	//user, err := client.CreateUser(ctx, &api.NewUser{Email: startemail})

	//log.Print(user)
	//log.Print(err)

	user, err := client.GetUserByEmail(ctx, &api.GetUserByEmailRequest{
		Email: newemail,
	})

	if err != nil {
		log.Fatalf("Failed to get user by email %v error %v", user, err)
	}

	newUser, err := client.UpdateUserEmail(ctx, &api.UpdateUserEmailRequest{
		UserId: user.Id,
		Email:  newemail,
	})

	log.Print(newUser)
	log.Print(err)

	client.RemoveUser(ctx, &api.RemoveUserRequest{UserId: user.Id})

	//deploymentId := "c1e2052c-06f3-4a2b-a8a5-73ac7651c022"
	//peripheralId := "51e6f06e-1a68-11eb-888c-0242ac110002"

	//for i := 0; i < 5; i++ {
	//	thetime, err := ptypes.TimestampProto(
	//		time.Date(2020, 10, 31, 19, 2*i, 0, 0, time.UTC),
	//	)
	//	_, err = client.SendEvent(
	//		ctx,
	//		&api.NewMeasurementEvent{
	//			PeripheralId: peripheralId,
	//			DeploymentId: deploymentId,
	//			Value:        80,
	//			TimeStamp:    thetime,
	//		},
	//	)

	//	log.Printf("Sent a request did it fail %v", err)
	//}
	//start, err := ptypes.TimestampProto(time.Date(2020, 10, 31, 19, 0, 0, 0, time.UTC))
	//end, err := ptypes.TimestampProto(time.Date(2020, 10, 31, 19, 7, 0, 0, time.UTC))
	//if err != nil {
	//	log.Fatalf("The times are bad, err %v", err)
	//}
	//eventStreamClient, streamError := client.FilterEvents(ctx, &api.MeasurementEventFilterRequest{
	//	PeripheralIds: []string{peripheralId},
	//	DeploymentId:  deploymentId,
	//	StartTime:     start,
	//	EndTime:       end,
	//})

	//if streamError != nil {
	//	log.Fatalf("Failed to start receiving events, error %v", streamError)
	//}

	//for {
	//	newEvent, err := eventStreamClient.Recv()
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
