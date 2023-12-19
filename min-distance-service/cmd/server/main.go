package main

import (
	"log"
	"net"

	"github.com/ulmk/go-grpc-service/min-distance-service/pb"
	"github.com/ulmk/go-grpc-service/min-distance-service/server"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMinDistanceServiceServer(s, &server.GrpcServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
