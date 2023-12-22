package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sxline/smpls-project/grpc-service/internal/constants"
	"github.com/sxline/smpls-project/grpc-service/internal/model"
)

// WriteRepository is an interface defining a method for writing data to Elasticsearch.
type WriteRepository interface {
	Write(data model.DataModel) error
}

// writeRepo is the implementation of the WriteRepository interface.
type writeRepo struct {
	client *elasticsearch.Client
}

// NewWriteRepository creates a new WriteRepository with the provided Elasticsearch client.
func NewWriteRepository(client *elasticsearch.Client) WriteRepository {
	return &writeRepo{client: client}
}

// Write writes a DataModel to Elasticsearch.
func (r *writeRepo) Write(data model.DataModel) error {
	// Serialize DataModel to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Create an Elasticsearch IndexRequest
	req := esapi.IndexRequest{
		Index:      constants.ElasticIndexData,
		DocumentID: data.Id,
		Body:       bytes.NewReader(jsonData),
	}

	// Execute the IndexRequest
	res, err := req.Do(context.Background(), r.client)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()

	// Check for errors in the response
	if res.IsError() {
		return err
	}

	// Print the HTTP status code (for illustration purposes)
	fmt.Println(res.StatusCode)

	return nil
}
