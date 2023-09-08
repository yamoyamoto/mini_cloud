package main

import (
	"context"
	"github.com/yamoyamoto/mini_cloud/client/internal/app"
	"github.com/yamoyamoto/mini_cloud/client/internal/log"
	"os"
)

func main() {
	ctx := log.WithContext(context.Background(), log.New())
	app := app.New()

	if err := app.Run(ctx); err != nil {
		log.FromContext(ctx).Error(err.Error())
		os.Exit(1)
	}
}
