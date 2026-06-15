package app

import (
	"io/fs"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/mayron1806/api-template/config"
	"github.com/mayron1806/api-template/web"
)

func NewFiber(cfg *config.Config) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		AppName: "API Template",
	})
	dist, _ := fs.Sub(web.Dist, "dist")

	app.Get("/*", static.New("", static.Config{
		FS: dist,
		// Optional: Configure caching, browsing, etc.
		Browse: false,
	}))

	return app, nil
}
