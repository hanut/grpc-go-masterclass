package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hanut/grpc-go-masterclass/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hi, I am a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldnt connect to the server: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Connection: %f", c)

	doUnary(&c)

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

	fmt.Printf("Response of invocation: %v\n", res.Result)
}
