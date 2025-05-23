package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	mainpb "grpc_stream_client/proto/gen"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := mainpb.NewCalculateClient(conn)

	ctx := context.Background()
	// -------- SERVER SIDE STREAMING STARTS
	req := &mainpb.FibonacciRequest{
		N: 10,
	}
	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil {
		log.Fatalln("Error calling GenerateFibonacci func:", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			break
		}
		if err != nil {
			log.Fatalln("Error receiving data from GenerateFibonacci func:", err)
		}
		log.Println("Fibonacci number: ", resp.GetNumber())
	}

	// CLIENT SIDE STREAMING

	stream1, err := client.SendNumbers(ctx)
	if err != nil {
		log.Fatalln("Error creating stream:", err)
	}

	for num := range 9 {
		log.Println("Sending number: ", num)
		err := stream1.Send(&mainpb.NumberRequest{Number: int32(num)})
		if err != nil {
			log.Fatalln("Error sending data:", err)
		}
		time.Sleep(time.Second)
	}

	res, err := stream1.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error receiving response:", err)
	}

	log.Println("Sum: ", res.GetSum())

	// BIDIRECTIONAL STREAMING

	chatStream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalln("Error creating chat stream:", err)
	}

	waitc := make(chan struct{})
	// Send messages in a goroutine
	go func() {
		messages := []string{"Hello", "How are you?", "Goodbye"}
		for _, message := range messages {
			log.Println("Sending message:", message)
			err := chatStream.Send(&mainpb.ChatMessage{Message: message})
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second)
		}
		chatStream.CloseSend()
	}()

	// Receive messages in goroutine
	go func() {
		for {
			res, err := chatStream.Recv()
			if err == io.EOF {
				log.Println("End of stream")
				break
			}
			if err != nil {
				log.Fatalln("Error receiving data from GenerateFibonacci func:", err)
			}
			log.Println("Received response: ", res.GetMessage())
		}
		close(waitc)
	}()
	<-waitc
}
