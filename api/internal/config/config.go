package config

type GrpcConfig struct {
	GRPCAddress string
}

func GetGrpcConfig() GrpcConfig {
	return GrpcConfig{
		GRPCAddress: "localhost:50051",
	}
}

type Config struct {
	HTTPAddr string
}

func GetConfig() Config {
	return Config{
		HTTPAddr: ":8000",
	}
}
