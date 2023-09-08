package main

import (
	"context"
	"github.com/yamoyamoto/mini_cloud/client/internal/app"
	"github.com/yamoyamoto/mini_cloud/client/internal/log"
	"os"
)

func main() {
	ctx := log.WithContext(context.Background(), log.New())
	a := app.New()

	if err := a.Run(ctx); err != nil {
		log.FromContext(ctx).Error(err.Error())
		os.Exit(1)
	}
}
