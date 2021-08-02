package main

import (
	"github.com/MichaelDarr/stack/auth/internal/config"
	"github.com/MichaelDarr/stack/auth/internal/server"
)

func main() {
	cfg := config.New()

	go server.GRPC(&cfg)

	select {}
}
