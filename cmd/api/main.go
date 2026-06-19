// @title           API Template
// @version         1.0
// @description     Documentation for the API Template.
// @BasePath        /api
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/usesnipet/go-template/config"
	_ "github.com/usesnipet/go-template/docs"
	"github.com/usesnipet/go-template/internal/bootstrap"
	"github.com/usesnipet/go-template/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	level, parseErr := logger.ParseLevel(cfg.Log.Level)
	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "warning: %v\n", parseErr)
		level = logger.LevelInfo
	}

	appLogger := logger.NewLogger(level)
	if parseErr != nil {
		appLogger.Warn(parseErr.Error())
	}

	bootstrap.Bootstrap(cfg, appLogger)
}
