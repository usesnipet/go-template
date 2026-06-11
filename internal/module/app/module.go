package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mayron1806/api-template/internal/logger"
	"github.com/mayron1806/api-template/internal/module/config"
	"github.com/mayron1806/api-template/internal/module/database"
	"github.com/mayron1806/api-template/internal/module/user"
	"go.uber.org/fx"
)

var Module = fx.Module("app",
	database.Module,
	user.Module,
	fx.Provide(fiber.New),
	fx.Invoke(func(
		app *fiber.App,
		cfg *config.Config,
		log *logger.Logger,
		lc fx.Lifecycle,
	) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Infof("starting server on port %d", cfg.Server.Port)
				if err := app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
					log.Errorf("failed to start server: %v", err)
				}
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return app.Shutdown()
			},
		})
	}),
)
