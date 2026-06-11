package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mayron1806/api-template/internal/module/config"
	"github.com/mayron1806/api-template/internal/module/database"
	"github.com/mayron1806/api-template/internal/module/user"
	"go.uber.org/fx"
)

var Module = fx.Module("app",
	config.Module,
	database.Module,
	user.Module,
	fx.Provide(fiber.New),
	fx.Invoke(func(
		app *fiber.App,
		cfg *config.Config,
		lc fx.Lifecycle,
	) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return app.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
			},
			OnStop: func(ctx context.Context) error {
				return app.Shutdown()
			},
		})
	}),
)
