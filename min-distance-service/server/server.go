package server

import (
	"context"

	"github.com/ulmk/go-grpc-service/min-distance-service/pb"
)

type GrpcServer struct {
	pb.UnimplementedMinDistanceServiceServer
}

// func (s *GrpcServer) GetMinDistance(ctx context.Context, req *MinDistanceRequest) (*MinDistanceResponse, error) {
func (s *GrpcServer) GetMinDistance(ctx context.Context, req *pb.MinDistanceRequest) (*pb.MinDistanceResponse, error) {
	nums := req.GetNums()

	pairs := minDistance(nums)

	resp := &pb.MinDistanceResponse{}

	for _, pair := range pairs {
		resp.Pairs = append(resp.Pairs, &pb.MinPair{
			Num1: pair[0],
			Num2: pair[1],
		})
	}

	return resp, nil
}
