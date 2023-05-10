package main

import (
	"log"
	"vk-intern-bot/config"
	"vk-intern-bot/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
