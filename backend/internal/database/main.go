package database

import (
	"github.com/MichaelDarr/shelf/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection is a configured database connection
type Connection struct {
	DB     *gorm.DB
	Config *config.PostgresConfig
}

// Open opens a GORM connection
func Open(cfg *config.PostgresConfig) (*Connection, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Connection{
		DB:     db,
		Config: cfg,
	}, nil
}
