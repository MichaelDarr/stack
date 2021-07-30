package main

import (
	"github.com/MichaelDarr/shelf/auth/internal/config"
	"github.com/MichaelDarr/shelf/auth/internal/server"
)

func main() {
	cfg := config.New()

	server.GRPC(&cfg)
}
