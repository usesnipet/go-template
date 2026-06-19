package bootstrap

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
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

	handler, apiRouter, serve, err := app.NewRouter(cfg)
	if err != nil {
		logger.Errorf("failed to create router: %v", err)
		return
	}

	apiRouter.Route("/users", func(r chi.Router) {
		userHandler.RegisterRoutes(r, serve)
	})

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Errorf("failed to listen on %s: %v", addr, err)
		return
	}

	server := &http.Server{Handler: handler}

	go func() {
		logger.Infof("server listening on %s", addr)
		if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
			logger.Errorf("server listener stopped: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
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
