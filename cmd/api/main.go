package main

import (
	"context"

	"github.com/mayron1806/api-template/internal/module/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		app.Module,
	).Start(context.Background())
}
