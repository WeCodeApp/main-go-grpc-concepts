package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	mainpb "grpc_stream_server/proto/gen"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	mainpb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *mainpb.AddRequest) (*mainpb.AddResponse, error) {
	return &mainpb.AddResponse{Sum: req.A + req.B}, nil
}

func (s *server) GenerateFibonacci(req *mainpb.FibonacciRequest, stream mainpb.Calculate_GenerateFibonacciServer) error {
	n := req.N
	a, b := 0, 1

	for i := 0; i < int(n); i++ {
		err := stream.Send(&mainpb.FibonacciResponse{
			Number: int32(a),
		})
		if err != nil {
			return err
		}
		log.Println("Sent number:", a)
		a, b = b, a+b
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) SendNumbers(stream mainpb.Calculate_SendNumbersServer) error {
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&mainpb.NumberResponse{Sum: sum})
		}
		if err != nil {
			return err
		}
		log.Println(req.GetNumber())
		sum += req.GetNumber()
	}
}

func (s *server) Chat(stream mainpb.Calculate_ChatServer) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Receiving values/messages from stream
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Received Message:", req.GetMessage())

		// Read input from the terminal
		fmt.Print("Enter response: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		input = strings.TrimSpace(input)

		response := &mainpb.ChatMessage{
			Message: input,
		}
		// Sending data/messages/values through the stream
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
	fmt.Println("Returning control")
	return nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	mainpb.RegisterCalculateServer(grpcServer, &server{})

	// Enable reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Println("Server started on port: ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
