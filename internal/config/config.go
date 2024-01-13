package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config структура конфигурации проекта
type Config struct {
	GRPC GRPCConfig
	HTTP HTTPConfig
	DB   PgConfig
}

func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DB.Host,
		c.DB.Port,
		c.DB.User,
		c.DB.Password,
		c.DB.Db,
	)
}

// MustConfig загружает конфигурацию из .env файла
func MustConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	return Config{
		GRPC: GRPCConfig{
			Port: getEnv("GRPC_PORT", "50051"),
			Host: getEnv("GRPC_HOST", "localhost"),
		},
		HTTP: HTTPConfig{
			Port: getEnv("HTTP_PORT", "8080"),
			Host: getEnv("HTTP_HOST", "localhost"),
		},
		DB: PgConfig{
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
