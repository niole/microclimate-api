package main

import (
	api "api/protobuf"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"rsc.io/quote"
	"time"
)

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
	in := &api.GetDeploymentRequest{
		OwnerUserId:  1,
		DeploymentId: 1,
	}
	deployment, err := client.GetDeployment(ctx, in)
	fmt.Println(deployment)
	fmt.Println(err)
}
