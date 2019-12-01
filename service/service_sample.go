package service

import (
	"context"

	"github.com/ybalexdp/sample-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type SampleService struct {
}

func (s *SampleService) GetSample(ctx context.Context, message *pb.GetSampleRequest) (*pb.SampleResponse, error) {
	if message.Name == "panic" {
		panic("failed")
		return nil, grpc.Errorf(codes.Internal, "Unexpected error")
	}
	return &pb.SampleResponse{
		Message: "Hello :" + message.Name,
	}, nil
}
