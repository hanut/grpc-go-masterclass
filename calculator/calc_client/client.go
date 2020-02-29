package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/hanut/grpc-go-masterclass/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	// args := os.Args
	// if len(args) < 2 {
	// 	args = append(args, "210")
	// }
	fmt.Println("CalculatorService client started...")
	// n, err := strconv.Atoi(args[1])

	// if err != nil {
	// 	log.Fatalf("Looks like the number you entered is invalid: %v", err)
	// }

	// fmt.Printf("RPC decompose for number: %v\n", n)

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error connecting to calculator service: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	doSum(&c)
	// doDecomposePN(&c, &n)
	// doCalculateAvg(&c, &([]uint32{245, 618, 718, 121}))

}

func doSum(c *calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		A: 3,
		B: 10,
	}

	res, err := (*c).Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error invoking Sum(): %v\n", err)
	}
	fmt.Printf("Result of Sum(): %v\n", res.GetResult())
}

func doDecomposePN(c *calculatorpb.CalculatorServiceClient, n *int) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: uint32(*n),
	}

	tmp := make([]uint32, 0)
	rStream, err := (*c).PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("Error invoking PrimeNumberDecomposition(): %v", err)
	}

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

func doCalculateAvg(c *calculatorpb.CalculatorServiceClient, nums *[]uint32) {
	stream, err := (*c).ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error opening stream in doCalculateAvg(): %v", err)
	}

	for _, n := range *nums {
		req := &calculatorpb.ComputeAverageRequest{
			Number: n,
		}
		stream.Send(req)
		time.Sleep(time.Millisecond * 500)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing stream: %v\n", err)
	}
	avg := res.GetResult()

	fmt.Printf("Average (%v) = %v\n", *nums, avg)
}
