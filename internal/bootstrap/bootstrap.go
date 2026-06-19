package bootstrap

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/usesnipet/go-template/config"
	"github.com/usesnipet/go-template/internal/logger"
	"github.com/usesnipet/go-template/internal/module/app"
	"github.com/usesnipet/go-template/internal/module/database"
	"github.com/usesnipet/go-template/internal/module/user"
)

func Bootstrap(cfg *config.Config, logger *logger.Logger) {
	//region Database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		logger.Errorf("failed to create database: %v", err)
		return
	}
	if err := database.RunMigrations(cfg, logger); err != nil {
		logger.Errorf("failed to run migrations: %v", err)
		return
	}
	//endregion

	//region Repositories
	userRepository := user.NewUserRepository(db, logger)
	//endregion

	//region Services
	userService := user.NewUserService(userRepository, logger)
	//endregion

	//region Handlers
	userHandler := user.NewUserHandler(userService, logger)
	//endregion

	fiberApp, router, err := app.NewFiber(cfg)
	if err != nil {
		logger.Errorf("failed to create fiber app: %v", err)
		return
	}

	userHandler.RegisterRoutes(router.Group("/users"))

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Errorf("failed to listen on %s: %v", addr, err)
		return
	}

	go func() {
		logger.Infof("server listening on %s", addr)
		if err := fiberApp.Listener(ln); err != nil {
			logger.Errorf("server listener stopped: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := fiberApp.ShutdownWithContext(ctx); err != nil {
		logger.Errorf("server shutdown failed: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("failed to get sql db: %v", err)
	} else if err := sqlDB.Close(); err != nil {
		logger.Errorf("failed to close database: %v", err)
	}

	logger.Info("server stopped")
}
