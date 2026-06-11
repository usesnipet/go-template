package config

import (
	"context"

	"github.com/joho/godotenv"
	envConfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server   ServerConfig   `env:", prefix=SERVER_"`
	Database DatabaseConfig `env:", prefix=DB_"`
}

func newConfig() *Config {
	_ = godotenv.Load()
	ctx := context.Background()
	var cfg = &Config{}
	var err = envConfig.Process(ctx, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
