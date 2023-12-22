package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/sxline/smpls-project/worker/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"sync"
	"testing"
	"time"
)

// MockWriteDataServiceClient is a mock for pb.WriteDataServiceClient for testing purposes.
type MockWriteDataServiceClient struct {
	WriteFunc  func(ctx context.Context, in *pb.Data, opts ...grpc.CallOption) (*emptypb.Empty, error)
	WriteCalls []struct {
		ctx context.Context
		in  *pb.Data
	}
}

// Write is the mock implementation of the Write method in pb.WriteDataServiceClient.
func (m *MockWriteDataServiceClient) Write(ctx context.Context, in *pb.Data, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	if m.WriteFunc != nil {
		m.WriteCalls = append(m.WriteCalls, struct {
			ctx context.Context
			in  *pb.Data
		}{ctx, in})
		return m.WriteFunc(ctx, in, opts...)
	}
	return nil, errors.New("not implemented")
}

// TestWorkerService_ReadDataFromChannel tests the readDataFromChannel method of the WorkerService.
func TestWorkerService_ReadDataFromChannel(t *testing.T) {
	mockClient := &MockWriteDataServiceClient{
		WriteFunc: func(ctx context.Context, in *pb.Data, opts ...grpc.CallOption) (*emptypb.Empty, error) {
			return &emptypb.Empty{}, nil
		},
	}
	workerService := &workerService{
		chanData:    make(chan pb.Data),
		wg:          &sync.WaitGroup{},
		writeClient: mockClient,
	}
	workerService.wg.Add(1)
	// Start the readDataFromChannel goroutine
	go workerService.readDataFromChannel()

	// Create test data
	testData := pb.Data{XId: "1", Type: "Test data"}

	// Send data to the channel
	workerService.chanData <- testData

	// Give the goroutine time to process the data
	time.Sleep(100 * time.Millisecond)

	// Close the channel
	close(workerService.chanData)

	workerService.wg.Wait()

	// Give the goroutine time to complete
	time.Sleep(100 * time.Millisecond)

	// Check that the Write mock method was called with the expected data
	assert.Equal(t, 1, len(mockClient.WriteCalls))
	call := mockClient.WriteCalls[0]
	assert.Equal(t, context.Background(), call.ctx)
	assert.Equal(t, &testData, call.in)
}

// TestWorkerService_ReadDataAndSendToChannel tests the readDataAndSendToChannel method of the WorkerService.
func TestWorkerService_ReadDataAndSendToChannel(t *testing.T) {
	workerService := &workerService{
		chanData: make(chan pb.Data),
		wg:       &sync.WaitGroup{},
	}

	workerService.wg.Add(2)
	// Start the readDataAndSendToChannel goroutine
	go workerService.readDataAndSendToChannel("../../data_test.json")

	var dataArr []*pb.Data
	mutex := &sync.Mutex{} // Create a mutex to protect access to dataArr

	go func() {
		defer workerService.wg.Done()
		for {
			data, ok := <-workerService.chanData
			if !ok {
				break
			}

			// Protect access to dataArr using the mutex
			mutex.Lock()
			dataArr = append(dataArr, &data)
			mutex.Unlock()
		}
	}()

	workerService.wg.Wait()

	// Check that the Write mock method was not called (as expected, since the channel was closed)
	assert.Equal(t, 2, len(dataArr))
}
