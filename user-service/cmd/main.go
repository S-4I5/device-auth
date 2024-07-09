package main

import (
	"context"
	"user-service/internal/app"
	"user-service/internal/config"
)

func main() {
	cfg := config.MustRead("./config/config.yaml")

	ctx := context.TODO()
	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		panic(err.Error())
	}

	if err = a.Start(); err != nil {
		panic(err.Error())
	}
}
