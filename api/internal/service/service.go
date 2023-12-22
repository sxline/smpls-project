package serv

import (
	"api/internal/pb"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DataService is an interface defining methods for fetching data and statistics
type DataService interface {
	GetAllData(searchText string, from, size int32) (*pb.GetAllDataResponse, error)
	GetStatistic() (*pb.StatisticResponse, error)
}

// dataServ is an implementation of the DataService interface.
type dataServ struct {
	grpcClient pb.ReadDataServiceClient
}

// NewDataService creates a new DataService instance with the provided ReadDataServiceClient
func NewDataService(grpcClient pb.ReadDataServiceClient) DataService {
	return &dataServ{grpcClient: grpcClient}
}

// GetAllData fetches data based on the provided search text, from, and size parameters
func (s dataServ) GetAllData(searchText string, from, size int32) (*pb.GetAllDataResponse, error) {
	return s.grpcClient.GetAllData(context.Background(), &pb.GetAllDataRequest{
		Text: searchText,
		From: from,
		Size: size,
	})
}

// GetStatistic fetches statistics
func (s dataServ) GetStatistic() (*pb.StatisticResponse, error) {
	return s.grpcClient.GetStatistic(context.Background(), &emptypb.Empty{})
}
