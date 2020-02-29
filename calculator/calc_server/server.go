package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hanut/grpc-go-masterclass/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	a := req.GetA()
	b := req.GetB()
	res := calculatorpb.SumResponse{
		Result: a + b,
	}
	return &res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition() was invoked with \n%v\n", req)
	n := req.GetNumber()
	k := uint32(2)
	for n > 1 {
		if n%k == 0 {
			n = n / k
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Response: k,
			}
			stream.Send(res)
		} else {
			k++
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	total := uint32(0)
	count := uint32(0)
	for {
		m, err := stream.Recv()
		if err == io.EOF {
			// Stream has ended, return average
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: float64(total) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v\n", err)
			return err
		}
		n := m.GetNumber()
		count++
		total = total + n
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Printf("FindMaximum invoked with %v\n", stream)

	max := int32(0)
	started := false

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Stream has ended
			return nil
		}
		if err != nil {
			fmt.Printf("Error reading stream: %v\n", err)
			return err
		}
		num := req.GetNum()
		fmt.Printf("Received number: %v\n", num)
		if !started {
			max = num
			started = true
		}
		if num > max {
			max = num
		}
		res := &calculatorpb.FindMaximumResponse{
			Max: max,
		}
		stream.Send(res)
	}
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	num := req.GetNumber()
	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v\n", num),
		)
	}
	return &calculatorpb.SquareRootResponse{
		Root: math.Sqrt(float64(num)),
	}, nil
}

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("Calculator Service listening on 0.0.0.0[:50051]")

	err = s.Serve(l)

	if err != nil {
		log.Fatalf("Error connecting to GRPC Service: %v", err)
	}

}
