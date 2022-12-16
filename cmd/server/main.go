package main

import (
	"context"
	"flag"

	"github.com/status-mok/server/internal/pkg/log"
	"github.com/status-mok/server/internal/server/app"
)

var flagConfigFile = flag.String("f", "", "path to configuration yaml file")

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := app.NewApp()

	if err := application.Run(ctx, *flagConfigFile); err != nil {
		log.L(ctx).Fatal(err)
	}
}
