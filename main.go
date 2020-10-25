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

	log.Println("Testing get deployment")
	newDeployment, _ := client.GetDeployment(ctx, &api.GetDeploymentRequest{
		OwnerUserId:  "e5f1f661-679c-4f6e-a2ab-6edbcf05a3df",
		DeploymentId: "819d0ce7-16eb-11eb-888c-0242ac110002",
	})

	fmt.Println(newDeployment)
}
