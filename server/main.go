package main

import (
	"context"
	"log"
	"net"

	pb "github.com/aleiphoenix/gopush/api/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Push(
	ctx context.Context, in *pb.PushRequest) (*pb.PushResponse, error) {

	return &pb.PushResponse{
		Status: 1,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPushServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
