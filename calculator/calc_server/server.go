package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hanut/grpc-go-masterclass/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

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
