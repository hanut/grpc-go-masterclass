package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/hanut/grpc-go-masterclass/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client started...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldnt connect to the server: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Connection: %f", c)

	doUnary(&c)

	doServerStreaming(&c)

}

func doUnary(c *greetpb.GreetServiceClient) {
	fmt.Println("Unary RPC call...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hanut",
			LastName:  "Singh Gusain",
		},
	}

	res, err := (*c).Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error invoking greet: %v", err)
	}

	fmt.Printf("Response of Greet: %v\n", res.Result)
}

func doServerStreaming(c *greetpb.GreetServiceClient) {
	fmt.Println("Server Streaming RPC call...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hanut",
			LastName:  "Singh Gusain",
		},
		Reps: 10,
	}

	rStream, err := (*c).GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to invoke GreetManyTimes: %v", err)
	}

	for {
		msg, err := rStream.Recv()
		if err == io.EOF {
			// Stream ended
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
			break
		}
		fmt.Printf("Response from GreetManyTimes: %v\n", msg)
	}

}
