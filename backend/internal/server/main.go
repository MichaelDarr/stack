package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MichaelDarr/shelf/backend/internal/config"
	"github.com/MichaelDarr/shelf/backend/internal/database"
)

// HTTP is the backend server
func HTTP(cfg *config.ServerConfig) {
	ctx := context.Background()

	db, err := database.Open(&cfg.Postgres)
	if err != nil {
		log.Fatalf("postgres connection failed: %v", err)
	}

	router := Route(cfg, db)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// create interrupt listener
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// serve application until a fatal error occurs
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("fatal server error: %s", err)
		}
	}()

	// lock releases when the shutoff signal is received, the
	<-done

	log.Printf("server execution terminated")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Printf("server exited properly")

	defer os.Exit(0)
}
