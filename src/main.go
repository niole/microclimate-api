package main

import (
	"fmt"
	pb "protobuf/tutorial/tutorialpb"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	fmt.Println(p)
}
