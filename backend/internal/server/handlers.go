package server

import (
	"net/http"

	"github.com/MichaelDarr/shelf/backend/internal/config"
	"github.com/MichaelDarr/shelf/backend/internal/database"
	"github.com/MichaelDarr/shelf/backend/internal/graphql/generated"
	"github.com/MichaelDarr/shelf/backend/internal/resolvers"
	ur "github.com/MichaelDarr/shelf/backend/internal/user/repository"
	us "github.com/MichaelDarr/shelf/backend/internal/user/service"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

// GraphQL is the primary graphql handler
func GraphQL(cfg *config.ServerConfig, connection *database.Connection) http.HandlerFunc {
	userRepo := ur.NewUserRepository(connection.DB)
	userSvc := us.NewUserService(userRepo)

	c := generated.Config{
		Resolvers: &resolvers.Resolver{
			Config:      cfg,
			UserService: userSvc,
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
