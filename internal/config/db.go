package config

type pgConfig struct {
	Host     string `env:"PG_HOST" envDefault:"localhost"`
	Port     string `env:"PG_PORT" envDefault:"5432"`
	User     string `env:"PG_USER" envDefault:"postgres"`
	Password string `env:"PG_PASSWORD" envDefault:"postgres"`
	Db       string `env:"PG_DB" envDefault:"postgres"`
}
