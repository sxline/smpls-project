package elasticsearch

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sxline/smpls-project/grpc-service/internal/config"
	"github.com/sxline/smpls-project/grpc-service/internal/constants"
	"log"
	"strings"
)

// ElasticsearchConnection initializes and returns an Elasticsearch client.
func ElasticsearchConnection(cfg *config.ElasticsearchConfig) *elasticsearch.Client {
	// Configure Elasticsearch client settings.
	elasticCfg := elasticsearch.Config{
		Addresses: []string{cfg.Address},
	}

	// Create a new Elasticsearch client.
	client, err := elasticsearch.NewClient(elasticCfg)
	if err != nil {
		log.Fatalf("elasticsearch connection failure: %v", err)
	}

	// Ping Elasticsearch to check the connection.
	res, err := client.Ping()
	if err != nil {
		log.Fatalf("elasticsearch ping failure: %v", err)
	}
	if res.IsError() {
		log.Fatalf("elasticsearch ping failure: %v", res)
	}

	// Check if the specified index exists.
	resp, err := esapi.IndicesExistsRequest{
		Index: []string{constants.ElasticIndexData},
	}.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("elasticsearch ping failure: %v", err)
	}
	if resp.StatusCode == 404 {
		// If the index does not exist, create it.
		if err = setIndex(client, constants.ElasticIndexData); err != nil {
			log.Fatalf("elasticsearch ping failure: %v", err)
		}
	}

	return client
}

// setIndex creates an Elasticsearch index with specified settings and mappings.
func setIndex(client *elasticsearch.Client, index string) error {
	// Define the index settings and mappings in JSON format.
	mapping := `
	{
	  "settings": {
		"analysis": {
		  "analyzer": {
			"ru_analyzer": {
			  "tokenizer": "standard",
			  "filter": [
				"lowercase",
				"ru_RU",
				"ru_stemmer"
			  ]
			},
			"ro_analyzer": {
			  "tokenizer": "standard",
			  "filter": [
				"lowercase",
				"ro_RO",
				"ro_stemmer"
			  ]
			}
			
		  },
		  "filter": {
			"ru_stemmer": {
			  "type": "stemmer",
			  "language": "russian"
			},
			"ru_RU": {
			  "type": "hunspell",
			  "locale": "ru_RU"
			},
			"ro_stemmer": {
			  "type": "stemmer",
			  "language": "romanian"
			},
			"ro_RO": {
			  "type": "hunspell",
			  "locale": "ro_RO"
			}
		  }
		}
	  },
	  "mappings": {
		"properties": {
		  "title.ru": {"type": "string", "analyzer": "ru_analyzer"},
		  "title.ro": {"type": "string", "analyzer": "ro_analyzer"}
		}
	  }
	}`

	// Create the Elasticsearch index.
	_, err := client.Indices.Create(
		index,
		client.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return err
	}
	return nil
}
