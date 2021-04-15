package main

import (
	"github.com/MichaelDarr/shelf/backend/internal/config"
	"github.com/MichaelDarr/shelf/backend/internal/server"
)

func main() {
	cfg := config.New()

	go server.HTTP(&cfg)

	select {}
}
