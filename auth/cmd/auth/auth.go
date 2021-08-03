package main

import (
	"log"

	"github.com/MichaelDarr/stack/auth/internal/config"
	"github.com/MichaelDarr/stack/auth/internal/server"
	"github.com/MichaelDarr/stack/auth/pkg/auth"
)

func main() {
	cfg := config.New()

	keyOne, err := auth.GenerateRandomRSAKey()
	if err != nil {
		log.Fatalf("Failed to generate auth first key: %v", err)
	}

	keyTwo, err := auth.GenerateRandomRSAKey()
	if err != nil {
		log.Fatalf("Failed to generate auth second key: %v", err)
	}

	keySet := auth.NewKeySet([]auth.Key{keyOne, keyTwo})

	grcpServer := server.NewGRPC(keySet)

	go grcpServer.Serve(cfg.Port)

	select {}
}
