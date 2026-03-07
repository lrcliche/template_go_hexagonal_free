package main

import (
	"context"
	"log"
	"os"

	"template-go-hexagonal/application/config"
	"template-go-hexagonal/presentation/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	cfg := config.MustLoad()

	httpServer, err := server.NewServer(context.Background(), cfg)
	if err != nil {
		log.Fatalf("[BOOT] failed to bootstrap server: %v", err)
	}

	if err := httpServer.Run(); err != nil {
		log.Fatalf("[BOOT] server execution failed: %v", err)
	}
}
