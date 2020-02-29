package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/hanut/grpc-go-masterclass/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	fname := req.GetGreeting().GetFirstName()
	lname := req.GetGreeting().GetLastName()
	res := greetpb.GreetResponse{
		Result: "Hello " + fname + " " + lname,
	}
	return &res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	fname := req.GetGreeting().GetFirstName()
	lname := req.GetGreeting().GetLastName()
	reps := req.GetReps()
	for i := uint32(0); i < reps; i++ {
		// Convert i to a string using strconv.FormatUint which takes a
		// uint64 so we also need to explicitly typecast to uint64
		iAsString := strconv.FormatUint(uint64(i), 10)

		result := "Hello " + fname + " " + lname + " (" + iAsString + ")"
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(300 * time.Millisecond)
	}
	return nil
}

func (*server) GreetLong(stream greetpb.GreetService_GreetLongServer) error {
	fmt.Printf("GreetLong function was invoked with %v\n", stream)

	res := []string{}
	count := uint32(0)

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			// Stream has ended
			break
		}
		if err != nil {
			log.Fatalf("Error processing string: %v", err)
			return err
		}
		count++
		fname := r.GetGreeting().GetFirstName()
		lname := r.GetGreeting().GetLastName()

		fmt.Printf("Received GreetLong Message for %v (%v)\n", fname+" "+lname, count)

		res = append(res, "Hello "+fname+" "+lname)
	}

	return stream.SendAndClose(&greetpb.GreetLongResponse{
		Result: res,
	})
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
