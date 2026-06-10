package config

type DatabaseConfig struct {
	Host     string `env:"HOST" default:"localhost"`
	Port     int    `env:"PORT" default:"5432"`
	User     string `env:"USER" default:"postgres"`
	Password string `env:"PASSWORD" default:"postgres"`
	Database string `env:"DATABASE" default:"postgres"`
	SSLMode  string `env:"SSL_MODE" default:"disable"`
	URL      string `env:"URL" default:"postgres://postgres:postgres@localhost:5432/postgres"`
}
