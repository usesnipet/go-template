package database

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(newDatabase),
	fx.Invoke(func(lc fx.Lifecycle, db *DB) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return db.Ping()
			},
			OnStop: func(ctx context.Context) error {
				return db.Close()
			},
		})
	}),
)
