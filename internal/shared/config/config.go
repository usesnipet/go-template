package config

import (
	"context"
	"sync"

	"github.com/joho/godotenv"
	envconfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server   ServerConfig   `env:", prefix=SERVER_"`
	Database DatabaseConfig `env:", prefix=DATABASE_"`
}

var once sync.Once
var cfg *Config
var err error

func Load() (*Config, error) {
	once.Do(func() {
		_ = godotenv.Load()
		ctx := context.Background()
		cfg = &Config{}
		err = envconfig.Process(ctx, cfg)
	})
	return cfg, err
}

func Get() *Config {
	if err != nil {
		panic(err)
	}
	if cfg == nil {
		panic("config not loaded")
	}
	return cfg
}
