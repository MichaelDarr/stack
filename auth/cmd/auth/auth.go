package main

import (
	"log"

	"github.com/MichaelDarr/stack/auth/internal/config"
	"github.com/MichaelDarr/stack/auth/internal/server"
	"github.com/MichaelDarr/stack/auth/pkg/auth"
)

func main() {
	cfg := config.New()

	key, err := auth.GenerateRandomKey()
	if err != nil {
		log.Fatalf("Failed to generate auth key: %v", err)
	}

	grcpServer := server.NewGRPC(key)

	go grcpServer.Serve(cfg.Port)

	select {}
}
