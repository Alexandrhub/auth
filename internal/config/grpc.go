package config

type grpcConfig struct {
	Port string `env:"GRPC_PORT" envDefault:"50051"`
	Host string `env:"GRPC_HOST" envDefault:"localhost"`
}
