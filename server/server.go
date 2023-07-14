package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "stressTesting/pb/dummy"
)

// server implements the gRPC server interface
type server struct {
	pb.UnimplementedDummyServiceServer
}

// YourMethod is an example RPC method
func (s *server) DummyMethod(ctx context.Context, request *pb.DummyRequest) (*pb.DummyResponse, error) {
	// Handle your logic here and return a response
	time.Sleep(20*time.Microsecond)
	response := & pb.DummyResponse{
		Message: "Hello, " + request.Name + "!",
	}
	return response, nil
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register your server implementation with the gRPC server
	pb.RegisterDummyServiceServer(s, &server{})

	// Start the gRPC server
	log.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}