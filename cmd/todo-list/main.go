package main

import (
	"log"
	"todo-list/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
}
