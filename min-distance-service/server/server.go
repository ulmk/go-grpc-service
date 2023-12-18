package server

import "context"

type server struct{}

func (s *server) GetMinDistance(ctx context.Context, req *MinDistanceRequest) (*MinDistanceResponse, error) {

	nums := req.GetNums()

	pairs := minDistance(nums)

	resp := &MinDistanceResponse{}

	for _, pair := range pairs {
		resp.Pairs = append(resp.Pairs, &MinPair{
			Num1: pair[0],
			Num2: pair[1],
		})
	}

	return resp, nil
}
