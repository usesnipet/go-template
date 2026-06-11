package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mayron1806/api-template/internal/logger"
	"github.com/mayron1806/api-template/internal/module/app"
	"github.com/mayron1806/api-template/internal/module/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
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

	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			return logger.NewFXEventLogger(appLogger)
		}),
		fx.Supply(cfg, appLogger),
		app.Module,
	).Start(context.Background())
}
