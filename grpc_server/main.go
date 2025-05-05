package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{Sum: req.A + req.B}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)

	grpcServer := grpc.NewServer()

	pb.RegisterCalculateServer(grpcServer, &server{})

	log.Println("Server started on port: ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
