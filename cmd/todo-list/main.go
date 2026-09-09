package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"todo-list/internal/api"
	"todo-list/internal/config"
	"todo-list/internal/postgres"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	pool, err := postgres.NewPool(ctx, &cfg.Postgres)
	if err != nil {
		log.Fatalf("pool creation: %v", err)
	}
	defer pool.Close()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.Http.Port),
		Handler: api.NewRouter(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()
	<-ctx.Done()

	shutDownCtx, shutDownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutDownCancel()

	if err := server.Shutdown(shutDownCtx); err != nil {
		log.Fatalf("server shutdown: %v", err)
	}
}
