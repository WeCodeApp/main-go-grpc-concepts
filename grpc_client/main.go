package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	mainpb "simplegrpcclient/proto/gen"
	"time"
)

func main() {

	cert := "cert.pem"
	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalf("failed to create creds: %v", err)
	}

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatal("Failed to connect", err)
	}
	defer conn.Close()

	client := mainpb.NewCalculateClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := mainpb.AddRequest{
		A: 10,
		B: 25,
	}

	res, err := client.Add(ctx, &req)
	if err != nil {
		log.Fatal("Failed to add", err)
	}
	log.Println("Result: ", res.Sum)
}
