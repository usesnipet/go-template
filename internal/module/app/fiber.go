package app

import (
	"io/fs"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/mayron1806/api-template/config"
	"github.com/mayron1806/api-template/web"
)

func NewFiber(cfg *config.Config) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		AppName: "API Template",
	})
	// Middlewares

	app.Get("/swagger/*", swaggo.HandlerDefault)

	// Serve the static files from the web/dist directory
	dist, _ := fs.Sub(web.Dist, "dist")
	app.Get("/*", static.New("", static.Config{
		FS: dist,
		// Optional: Configure caching, browsing, etc.
		Browse: false,
	}))

	return app, nil
}
