package main

import (
	"context"
	"fmt"
	"learnai/internal/adapters/config"
	"learnai/internal/adapters/storage/postgres"
	"log/slog"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()
	db, err := postgres.New(ctx, *config.DB)
	if err != nil {
		slog.Error("Error connecting to the database", "error", err)
		os.Exit(1)
	}

	defer db.Close()
}
