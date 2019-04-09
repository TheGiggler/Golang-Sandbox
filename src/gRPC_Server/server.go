package main

import (
	"context"
	pb "financial"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) ProcessTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	//log.Printf("Received: %v", in.Name)
	return &pb.TransactionResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFinancialTransactionServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("I'm listening!")
}
