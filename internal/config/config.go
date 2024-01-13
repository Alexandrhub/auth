package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Grpc grpcConfig
	Http httpConfig
	Db   pgConfig
}

func MustConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	return Config{
		Grpc: grpcConfig{
			Port: getEnv("GRPC_PORT", "50051"),
			Host: getEnv("GRPC_HOST", "localhost"),
		},
		Http: httpConfig{
			Port: getEnv("HTTP_PORT", "8080"),
			Host: getEnv("HTTP_HOST", "localhost"),
		},
		Db: pgConfig{
			Host:     getEnv("PG_HOST", "localhost"),
			Port:     getEnv("PG_PORT", "5432"),
			User:     getEnv("PG_USER", "postgres"),
			Password: getEnv("PG_PASSWORD", "postgres"),
			Db:       getEnv("PG_DB", "postgres"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
