package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/hanut/grpc-go-masterclass/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		args = append(args, "210")
	}
	fmt.Println("CalculatorService client started...")
	n, err := strconv.Atoi(args[1])

	if err != nil {
		log.Fatalf("Looks like the number you entered is invalid: %v", err)
	}

	fmt.Printf("RPC decompose for number: %v\n", n)

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error connecting to calculator service: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: uint32(n),
	}

	tmp := make([]uint32, 0)
	rStream, err := c.PrimeNumberDecomposition(context.Background(), req)

	for {
		res, err := rStream.Recv()

		if err == io.EOF {
			// Stream ended
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
		}
		n := res.GetResponse()
		fmt.Printf("Received: %v\n", n)
		tmp = append(tmp, n)
	}

	fmt.Printf("Decomposed number is: %v\n", tmp)
}
