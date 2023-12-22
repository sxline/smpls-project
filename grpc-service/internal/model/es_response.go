package model

// ElasticGetAllResponse represents the response structure for querying all documents from Elasticsearch
type ElasticGetAllResponse struct {
	Hits struct {
		Total struct {
			Value int32 `json:"value"`
		} `json:"total"`

		Hits []Hits `json:"hits"`
	} `json:"hits"`
}

// Hits represents the hits structure in the Elasticsearch response
type Hits struct {
	Index string  `json:"_index"`
	ID    string  `json:"_id"`
	Score float32 `json:"_score"`

	Source DataModel `json:"_source"`
}

// ElasticStatisticResponse represents the response structure for statistical aggregations from Elasticsearch
type ElasticStatisticResponse struct {
	Aggregations struct {
		SubcategoryAggregation struct {
			Buckets []struct {
				Key      string `json:"key"`
				DocCount int32  `json:"doc_count"`
			} `json:"buckets"`
		} `json:"subcategory_aggregation"`
	} `json:"aggregations"`
}
