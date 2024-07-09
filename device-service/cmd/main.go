package main

import (
	"device-service/internal/app"
	"device-service/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg := config.MustRead("./config/config.yaml")

	fmt.Println(cfg)

	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal("cannot setup server:", err.Error())
	}

	if err = a.Start(); err != nil {
		log.Fatal("stop server:", err.Error())
	}
}
