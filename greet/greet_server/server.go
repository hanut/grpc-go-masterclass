package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hanut/grpc-go-masterclass/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	fname := req.GetGreeting().GetFirstName()
	lname := req.GetGreeting().GetLastName()
	res := greetpb.GreetResponse{
		Result: "Hello " + fname + " " + lname,
	}
	return &res, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Server is listening on 0.0.0.0[:50051]")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
