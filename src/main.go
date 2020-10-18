package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	pb "protobuf/tutorial/tutorialpb"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	p := &pb.AddressBook{
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

	fmt.Println(p)

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fname := "/tmp/cat.txt"
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
