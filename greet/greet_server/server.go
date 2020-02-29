package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("Received a GreetEveryone invocation: %v\n", stream)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Stream ended, do cleanup
			return nil
		}
		if err != nil {
			fmt.Printf("Error reading from stream: %v\n", err)
			return err
		}
		fname := req.GetGreeting().GetFirstName()
		lname := req.GetGreeting().GetLastName()
		result := "Hello " + fname + " " + lname + "!"
		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalf("Error sending message to client: %v\n", err)
			return err
		}
	}
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline function was invoked with %v\n", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The client cancelled the request")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}
		time.Sleep(time.Second * 1)
	}
	fname := req.GetGreeting().GetFirstName()
	lname := req.GetGreeting().GetLastName()
	res := greetpb.GreetWithDeadlineResponse{
		Result: "Hello " + fname + " " + lname,
	}
	return &res, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	certFile := "/home/hanut/go/src/github.com/hanut/grpc-go-masterclass/ssl/localhost.crt"
	keyFile := "/home/hanut/go/src/github.com/hanut/grpc-go-masterclass/ssl/localhost.key"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Error creating credentials: %v\n", err)
		return
	}
	s := grpc.NewServer(grpc.Creds(creds))
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Server is listening on 0.0.0.0[:50051]")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
