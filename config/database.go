package config

type DatabaseConfig struct {
	URL          string `env:"URL"`
	Host         string `env:"HOST, default=localhost"`
	Port         int    `env:"PORT, default=5432"`
	User         string `env:"USER, default=postgres"`
	Password     string `env:"PASSWORD, default=postgres"`
	Database     string `env:"DATABASE, default=postgres"`
	SSLMode      string `env:"SSL_MODE, default=disable"`
	MaxOpenConns int    `env:"MAX_OPEN_CONNS, default=25"`
	MaxIdleConns int    `env:"MAX_IDLE_CONNS, default=5"`
}
