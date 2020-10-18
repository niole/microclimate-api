package main

import (
	pb "api/protobuf/tutorial/tutorialpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"rsc.io/quote"
	"time"
)

func main() {
	fmt.Println(quote.Go())
	addressbook := &pb.AddressBook{
		People: []*pb.Person{{
			Id:    1234,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-4321", Type: pb.Person_HOME},
			},
		},
		},
	}

	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	serverAddr := fmt.Sprintf("localhost:%d", 6000)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println(cancel)
	person, err := client.Search(ctx, addressbook)
	fmt.Println(person)
	fmt.Println(err)
}
