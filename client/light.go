package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "stressTesting/pb/dummy"
)

func main() {
	pid := os.Getpid()
	log.Println("PID:", pid)

	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new client instance
	client := pb.NewDummyServiceClient(conn)

	// Create a request message
	request := &pb.DummyRequest{
		Name: "John",
	}

    ticker := time.NewTicker(time.Second)
	interrupt := make(chan os.Signal, 1)

	for{
		select {
        case <-interrupt:
            return
		case <-ticker.C:
            // log.Println("Tick at", t)
			// Call the RPC method on the server
			for i:=0; i<10; i++ {
				_, err := client.DummyMethod(context.Background(), request)
				if err != nil {
					log.Fatalf("RPC failed: %v", err)
				}
			}
			// Print the response
			// log.Println("Done with all grpc calls")
    	}
	}

	// time.Sleep(time.Hour)
}
