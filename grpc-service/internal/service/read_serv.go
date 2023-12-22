package service

import (
	"context"
	"github.com/sxline/smpls-project/grpc-service/internal/model"
	"github.com/sxline/smpls-project/grpc-service/internal/pb"
	"github.com/sxline/smpls-project/grpc-service/internal/repo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type readServ struct {
	repo repo.ReadRepository
}

func NewReadService(repo repo.ReadRepository) pb.ReadDataServiceServer {
	return &readServ{repo: repo}
}

func (s readServ) GetAllData(_ context.Context, req *pb.GetAllDataRequest) (*pb.GetAllDataResponse, error) {
	resDataArr, total, err := s.repo.GetAll(req.GetText(), int(req.GetFrom()), int(req.GetSize()))
	if err != nil {
		return &pb.GetAllDataResponse{}, err
	}
	response := make([]*pb.Data, 0)
	for _, v := range resDataArr {
		response = append(response, model.ToGrpcData(v))
	}

	return &pb.GetAllDataResponse{
		Data:  response,
		Total: total,
	}, err

}

func (s readServ) GetStatistic(_ context.Context, _ *emptypb.Empty) (*pb.StatisticResponse, error) {
	resMap, err := s.repo.GetStatistic()
	if err != nil {
		return &pb.StatisticResponse{}, err
	}
	return &pb.StatisticResponse{Categories: resMap}, nil
}
