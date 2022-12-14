package main

import (
	"context"

	"github.com/status-mok/server/internal/pkg/log"
	"github.com/status-mok/server/internal/server/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := app.NewApp()

	if err := application.Start(ctx, ""); err != nil {
		log.L(ctx).Panic(err)
	}
}
