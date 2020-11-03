package controller

import (
	"context"
	pb "mastiff/proto"
)

type MastiffService struct{}

// 创建ID
func (s MastiffService) CreateId(ctx context.Context, req *pb.CreateIdReq) (*pb.CreateIdRes, error) {
	return &pb.CreateIdRes{
		Data: req.Prefix + "123456",
	}, nil
}
