package config

type Config struct {
	Port string
}

func GetConfig() *Config {
	return &Config{
		Port: "50051",
	}
}

type ElasticsearchConfig struct {
	Address string
}

func GetElasticsearchConfig() *ElasticsearchConfig {
	return &ElasticsearchConfig{
		Address: "http://localhost:9200",
	}
}
