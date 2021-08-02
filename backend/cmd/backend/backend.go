package main

import (
	"github.com/MichaelDarr/stack/backend/internal/config"
	"github.com/MichaelDarr/stack/backend/internal/server"
)

func main() {
	cfg := config.New()

	go server.HTTP(&cfg)

	select {}
}
