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
	fmt.Println(quote.Go())
	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	serverAddr := fmt.Sprintf("localhost:%d", 6000)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewDeploymentManagementServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Testing create deployment")

	newDeployment, err := client.CreateDeployment(ctx, &api.NewDeployment{
		OwnerUserId: newUuidString(),
	})
	log.Println(newDeployment.Id)
	log.Println(err)

	log.Println("Testing get deployment")
	gottenDeployment, gottenDeploymentErr := client.GetDeployment(ctx, &api.GetDeploymentRequest{
		OwnerUserId:  newDeployment.OwnerUserId,
		DeploymentId: newDeployment.Id,
	})
	log.Print(gottenDeployment.Id)
	log.Println(gottenDeploymentErr)

	log.Println("Testing delete deployment")
	_, err = client.RemoveDeployment(ctx, &api.RemoveDeploymentRequest{
		DeploymentId: newDeployment.Id,
	})

	log.Println(err)
}
