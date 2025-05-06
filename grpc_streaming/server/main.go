package main

import (
	"google.golang.org/grpc"
	mainpb "grpc_stream_server/proto/gen"
	"io"
	"log"
	"net"
	"time"
)

type server struct {
	mainpb.UnimplementedCalculateServer
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

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	mainpb.RegisterCalculateServer(grpcServer, &server{})

	log.Println("Server started on port: ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
