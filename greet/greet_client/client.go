package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	// doUnary(&c)

	// doServerStreaming(&c)

	// doClientStreaming(&c)

	doBidiStreaming(&c)
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

func doClientStreaming(c *greetpb.GreetServiceClient) {
	fmt.Println("Client streaming RPC call...")

	stream, err := (*c).GreetLong(context.Background())
	if err != nil {
		log.Fatalf("Error invoking GreetLong(): %v\n", err)
	}

	peeps := []*greetpb.GreetLongRequest{
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Hanut",
				LastName:  "Singh Gusain",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Aditya",
				LastName:  "Guleria",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Aniruddha",
				LastName:  "Mishra",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rishabh",
				LastName:  "Singh",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ishan",
				LastName:  "Bisht",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Shikhar",
				LastName:  "Srivastava",
			},
		},
	}

	for _, peep := range peeps {
		tmp := peep.GetGreeting()
		fmt.Printf("Sending message for %v\n", tmp)
		stream.Send(peep)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}

	fmt.Printf("Response from GreetLong(): %v", res)

}

func doBidiStreaming(c *greetpb.GreetServiceClient) {
	fmt.Println("BiDi streaming RPC call")

	stream, err := (*c).GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error opening stream: %v\n", err)
	}

	peeps := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Hanut",
				LastName:  "Singh Gusain",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Aditya",
				LastName:  "Guleria",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Aniruddha",
				LastName:  "Mishra",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rishabh",
				LastName:  "Singh",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ishan",
				LastName:  "Bisht",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Shikhar",
				LastName:  "Srivastava",
			},
		},
	}

	receiver := make(chan string)

	// Send a bunch of messages
	go func() {
		for i, peep := range peeps {
			fmt.Printf("Sending GreetEveryoneRequest for: %v\n", peep)
			err := stream.Send(peep)
			if err != nil {
				log.Fatalf("Error sending message for (%v): %v\n", peep, err)
			}
			if i < len(peeps)-1 {
				time.Sleep(time.Second)
			}
		}
		stream.CloseSend()
	}()

	// Receive a bunch on messages
	go func() {
		for {
			m, err := stream.Recv()
			if err == io.EOF {
				// Stream ended
				break
			}
			if err != nil {
				log.Fatalf("Error reading from stream: %v", err)
				break
			}
			// fmt.Println("Message from Service: " + m.GetResult())
			receiver <- m.GetResult()
		}
	}()

	for range peeps {
		msg := <-receiver
		fmt.Println(msg)
	}
	close(receiver)
}
