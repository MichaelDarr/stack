package server

import (
	"net/http"

	"github.com/MichaelDarr/shelf/backend/internal/config"
	"github.com/MichaelDarr/shelf/backend/internal/graphql/generated"
	"github.com/MichaelDarr/shelf/backend/internal/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

// GraphQL is the primary graphql handler
func GraphQL(cfg *config.ServerConfig) http.HandlerFunc {

	c := generated.Config{
		Resolvers: &resolvers.Resolver{
			Config: cfg,
		},
	}

	h := handler.New(generated.NewExecutableSchema(c))

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})
	h.Use(extension.Introspection{})

	return h.ServeHTTP
}

// Playground is the GraphQL playground handler
func Playground(cfg *config.ServerConfig) http.HandlerFunc {
	return playground.Handler("GraphQL", cfg.GQLPath)
}
