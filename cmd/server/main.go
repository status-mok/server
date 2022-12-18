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

	application, err := app.NewApp(ctx, *flagConfigFile)
	if err != nil {
		log.L(ctx).Fatal(err)
	}

	if err = application.Run(ctx); err != nil {
		log.L(ctx).Fatal(err)
	}
}
