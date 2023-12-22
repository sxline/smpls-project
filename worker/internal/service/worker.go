package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sxline/smpls-project/worker/internal/config"
	"github.com/sxline/smpls-project/worker/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"sync"
)

// WorkerService represents the interface for the worker service.
type WorkerService interface {
	Run(filename string)
}

// workerService is the implementation of the WorkerService interface.
type workerService struct {
	chanData    chan pb.Data              // Channel for communicating data between goroutines
	wg          *sync.WaitGroup           // WaitGroup for synchronizing goroutines
	writeClient pb.WriteDataServiceClient // gRPC client for writing data
}

// NewWorkerService creates a new instance of WorkerService.
func NewWorkerService() WorkerService {
	cfg := config.GetGrpcConfig()

	// Establish a gRPC connection
	conn, err := grpc.Dial(cfg.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print(err)
		return nil
	}

	client := pb.NewWriteDataServiceClient(conn)

	return &workerService{
		chanData:    make(chan pb.Data),
		wg:          &sync.WaitGroup{},
		writeClient: client,
	}
}

// Run starts the worker service.
func (s *workerService) Run(filename string) {
	fmt.Println("worker started")

	// Increment WaitGroup counter for each goroutine
	s.wg.Add(2)

	// Start goroutines for reading data from the channel and sending data to the channel
	go s.readDataFromChannel()
	go s.readDataAndSendToChannel(filename)

	// Wait for all goroutines to finish
	s.wg.Wait()

	fmt.Println("worker finished")
}

// readDataFromChannel reads data from the channel and sends it to the gRPC service.
func (s *workerService) readDataFromChannel() {
	// Decrement WaitGroup counter when the function exits
	defer s.wg.Done()

	for {
		data, ok := <-s.chanData
		if !ok {
			break
		}

		// Send data to the gRPC service
		_, err := s.writeClient.Write(context.Background(), &data)
		if err != nil {
			log.Print(err)
			return
		}
	}
}

// readDataAndSendToChannel reads data from a JSON file and sends it to the channel.
func (s *workerService) readDataAndSendToChannel(filePath string) {
	// Decrement WaitGroup counter when the function exits
	defer s.wg.Done()

	// Close the channel when the function exits
	defer close(s.chanData)

	file, err := os.Open(filePath)
	if err != nil {
		log.Print("Error opening file", err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Error closing file: %v", cerr)
		}
	}()

	decoder := json.NewDecoder(file)

	// Read opening delimiter. `[` or `{`
	if _, err = decoder.Token(); err != nil {
		log.Print("Error reading opening delimiter", err)
		return
	}

	for decoder.More() {
		var data pb.Data
		if err = decoder.Decode(&data); err != nil {
			log.Print("Error decoding file", err)
			return
		}

		// Send data to the channel
		s.chanData <- data
	}

	// Read closing delimiter. `]` or `}`
	if _, err = decoder.Token(); err != nil {
		log.Print("Error reading closing delimiter", err)
		return
	}
}
