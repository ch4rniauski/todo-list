package main

import (
	"context"
	"log"
	"os/signal"
	"todo-list/internal/config"
	"todo-list/internal/postgres"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	pool, err = postgres.NewPool(ctx, &cfg.Postgres)
	if err != nil {
		log.Fatalf("pool creation: %v", err)
	}
}
