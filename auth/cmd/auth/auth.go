package main

import (
	"log"

	"github.com/MichaelDarr/stack/auth/internal/config"
	"github.com/MichaelDarr/stack/auth/internal/server"
	"github.com/MichaelDarr/stack/auth/pkg/auth"
)

func main() {
	cfg := config.New()

	key, err := auth.GenerateRandomRSAKey()
	if err != nil {
		log.Fatalf("Failed to generate auth key: %v", err)
	}

	keySet := auth.NewKeySet([]auth.Key{key})

	grcpServer := server.NewGRPC(keySet)

	go grcpServer.Serve(cfg.Port)

	select {}
}
