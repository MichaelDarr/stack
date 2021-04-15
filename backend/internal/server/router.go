package server

import (
	"net/http"

	"github.com/MichaelDarr/shelf/backend/internal/config"
	"github.com/rs/cors"
)

// Route mounts all backend routes
func Route(cfg *config.ServerConfig) http.Handler {
	rtr := http.NewServeMux()

	rtr.Handle(cfg.GQLPath, GraphQL(cfg))
	rtr.Handle(cfg.PlaygroundPath, Playground(cfg))

	return cors.AllowAll().Handler(rtr)
}
