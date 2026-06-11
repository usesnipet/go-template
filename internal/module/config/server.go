package config

type ServerConfig struct {
	Port int `env:"PORT, default=8852"`
}
