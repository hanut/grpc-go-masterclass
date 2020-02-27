package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hanut/grpc-go-masterclass/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
