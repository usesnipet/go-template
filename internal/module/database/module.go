package database

import (
	"context"
	"fmt"

	"github.com/mayron1806/api-template/config"
	"github.com/mayron1806/api-template/internal/logger"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(newDatabase),
	fx.Invoke(func(
		lc fx.Lifecycle,
		db *DB,
		cfg *config.Config,
		logger *logger.Logger,
	) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				if err := db.Ping(); err != nil {
					return fmt.Errorf("ping database: %w", err)
				}
				if err := runMigrations(cfg, logger); err != nil {
					return fmt.Errorf("run migrations: %w", err)
				}
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return db.Close()
			},
		})
	}),
)
