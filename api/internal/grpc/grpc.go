package grpc

import (
	"api/internal/config"
	"api/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGrpcReadClient creates a new gRPC client for the ReadDataService
func NewGrpcReadClient(cfg config.GrpcConfig) (pb.ReadDataServiceClient, error) {
	// Establish a gRPC connection using the provided address
	conn, err := grpc.Dial(cfg.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// Return an error if the connection cannot be established.
		return nil, err
	}

	// Create a ReadDataServiceClient using the gRPC connection.
	client := pb.NewReadDataServiceClient(conn)

	// Return the gRPC client
	return client, nil
}
