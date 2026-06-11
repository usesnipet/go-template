package user

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Module("user",
	fx.Provide(newUserService, newUserHandler),
	fx.Invoke(func(app *fiber.App, handler *UserHandler) {
		handler.RegisterRoutes(app.Group("/users"))
	}),
)
