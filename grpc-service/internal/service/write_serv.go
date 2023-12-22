package service

import (
	"context"
	"github.com/sxline/smpls-project/grpc-service/internal/model"
	"github.com/sxline/smpls-project/grpc-service/internal/pb"
	"github.com/sxline/smpls-project/grpc-service/internal/repo"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewWriteService(repo repo.WriteRepository) pb.WriteDataServiceServer {
	return &write{repo: repo}
}

type write struct {
	repo repo.WriteRepository
}

func (s *write) Write(_ context.Context, req *pb.Data) (*emptypb.Empty, error) {
	data := model.GrpcDataToModel(req)
	if err := s.repo.Write(data); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
