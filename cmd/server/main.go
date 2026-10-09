package main

import (
	"context"
	"log"
	"time"

	"github.com/BurstWhite/goodnote-server/internal/database"
	"github.com/BurstWhite/goodnote-server/internal/note"
	"github.com/BurstWhite/goodnote-server/internal/router"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.Open(ctx, "goodnote.db")
	if err != nil {
		return err
	}
	defer db.Close()

	repository := note.NewRepository(db)
	service := note.NewService(repository)
	handler := note.NewHandler(service)
	r := router.NewRouter(handler)

	return r.Run(":8080")
}
