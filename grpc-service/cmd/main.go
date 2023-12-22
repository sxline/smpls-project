package main

import (
	"github.com/sxline/smpls-project/grpc-service/internal/config"
	"github.com/sxline/smpls-project/grpc-service/internal/pb"
	"github.com/sxline/smpls-project/grpc-service/internal/repo"
	"github.com/sxline/smpls-project/grpc-service/internal/service"
	"google.golang.org/grpc"

	esdb "github.com/sxline/smpls-project/grpc-service/internal/elasticsearch"

	"log"
	"net"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func run() error {
	cfg := config.GetConfig()
	elasticCfg := config.GetElasticsearchConfig()

	client := esdb.ElasticsearchConnection(elasticCfg)

	writeRepo := repo.NewWriteRepository(client)
	writeServ := service.NewWriteService(writeRepo)

	readRepo := repo.NewReadRepository(client)
	readServ := service.NewReadService(readRepo)

	listen, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", cfg.Port, err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterWriteDataServiceServer(grpcServer, writeServ)
	pb.RegisterReadDataServiceServer(grpcServer, readServ)

	log.Printf("gRPC server running on port %s", cfg.Port)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port %s: %v", cfg.Port, err)
	}
	return nil
}
