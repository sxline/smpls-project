package config

// GrpcConfig represents the configuration for gRPC settings.
type GrpcConfig struct {
	GRPCAddress string // GRPCAddress is the address where the gRPC server is running.
}

// GetGrpcConfig returns a default configuration for gRPC settings.
func GetGrpcConfig() GrpcConfig {
	return GrpcConfig{
		GRPCAddress: "localhost:50051", // Default gRPC server address.
	}
}
