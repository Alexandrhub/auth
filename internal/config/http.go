package config

type HTTPConfig struct {
	Port string `env:"HTTP_PORT" envDefault:"8080"`
	Host string `env:"HTTP_HOST" envDefault:"localhost"`
}
