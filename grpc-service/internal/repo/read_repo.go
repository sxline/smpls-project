package repo

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sxline/smpls-project/grpc-service/internal/constants"
	"github.com/sxline/smpls-project/grpc-service/internal/model"
	"strings"
)

// ReadRepository is an interface defining methods for reading data from Elasticsearch
type ReadRepository interface {
	GetAll(textSearch string, from int, size int) ([]model.DataModel, int32, error)
	GetStatistic() (map[string]int32, error)
}

// readRepo is the implementation of the ReadRepository interface.
type readRepo struct {
	client *elasticsearch.Client
}

// NewReadRepository creates a new ReadRepository with the provided Elasticsearch client.
func NewReadRepository(client *elasticsearch.Client) ReadRepository {
	return &readRepo{client: client}
}

// GetAll retrieves data from Elasticsearch based on the specified search parameters.
func (r readRepo) GetAll(textSearch string, from int, size int) ([]model.DataModel, int32, error) {
	query := r.generateGetAllQuery(textSearch)

	// Execute the search query
	res, err := r.client.Search(
		r.client.Search.WithIndex(constants.ElasticIndexData),
		r.client.Search.WithBody(strings.NewReader(query)),
		r.client.Search.WithFrom(from),
		r.client.Search.WithSize(size),
	)
	if err != nil {
		return nil, 0, err
	}

	// Close the response body when done
	defer func() { _ = res.Body.Close() }()

	// Decode the Elasticsearch response into a struct
	var resp model.ElasticGetAllResponse
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		fmt.Println("searchData", err)
		return nil, 0, err
	}

	// Process and return the response
	return r.processGetAllResponse(&resp)
}

// GetStatistic retrieves aggregated statistics from Elasticsearch.
func (r readRepo) GetStatistic() (map[string]int32, error) {
	query := r.generateGetStatisticQuery()

	// Execute the search query
	res, err := r.client.Search(
		r.client.Search.WithIndex(constants.ElasticIndexData),
		r.client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}

	// Close the response body when done
	defer func() { _ = res.Body.Close() }()

	// Decode the Elasticsearch response into a struct
	var resp model.ElasticStatisticResponse
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	// Process and return the response
	return r.processGetStatisticResponse(&resp)
}

// generateGetAllQuery generates the Elasticsearch query for GetAll based on the textSearch parameter.
func (r readRepo) generateGetAllQuery(textSearch string) string {
	if textSearch == "" {
		return `{
			"query": {
				"match_all": {}
			}
		}`
	}

	return fmt.Sprintf(`
    {
		"query": {
			"multi_match": {
				"query":    "%s",
				"fields": [ "title.ro", "title.ru" ], 
				"type":     "most_fields"
			}
		}
	}`, textSearch)
}

// generateGetStatisticQuery generates the Elasticsearch query for GetStatistic.
func (r readRepo) generateGetStatisticQuery() string {
	return `{
	  "size": 0,
	  "aggs": {
		"subcategory_aggregation": {
		  "terms": {
			"field": "categories.subcategory.keyword"
		  }
		}
	  }
	}`
}

// processGetAllResponse processes the Elasticsearch response for GetAll.
func (r readRepo) processGetAllResponse(resp *model.ElasticGetAllResponse) ([]model.DataModel, int32, error) {
	response := make([]model.DataModel, 0)
	for _, hit := range resp.Hits.Hits {
		data := hit.Source
		data.Id = hit.ID
		response = append(response, data)
	}
	return response, resp.Hits.Total.Value, nil
}

// processGetStatisticResponse processes the Elasticsearch response for GetStatistic.
func (r readRepo) processGetStatisticResponse(resp *model.ElasticStatisticResponse) (map[string]int32, error) {
	response := make(map[string]int32)
	for _, bucket := range resp.Aggregations.SubcategoryAggregation.Buckets {
		response[bucket.Key] = bucket.DocCount
	}
	return response, nil
}
