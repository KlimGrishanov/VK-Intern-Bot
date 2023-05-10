package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"vk-intern-bot/config"
	"vk-intern-bot/internal/repository"
	"vk-intern-bot/internal/services"
)

func Run(cfg *config.Config) {

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgx.Connect: %w", err))
	}
	defer func() {
		err = db.Close(context.Background())
		if err != nil {
			log.Fatal(fmt.Errorf("app - Run - db.Close: %w", err))
		}
	}()

	repo := repository.NewRepo(
		repository.NewUsersDataRepo(db),
	)

	handler, err := services.NewHandler(cfg.Bot.Token, repo)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - bot.NewHandler: %w", err))
	}
	err = handler.Handle()
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - handler.Handle: %w", err))
	}
}
