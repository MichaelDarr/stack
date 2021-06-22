package main

import (
	"fmt"

	"github.com/MichaelDarr/shelf/backend/internal/config"
	"github.com/MichaelDarr/shelf/backend/internal/migrate"
	"github.com/spf13/cobra"
)

const defaultMigrationName = "new_migration"

var createCmd = &cobra.Command{
	Use:   "create [migration name]",
	Short: "Create a migration",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		m, err := migrate.New(cfg.Postgres.MigrationFilepath, cfg.Postgres.URL())
		if err != nil {
			fmt.Printf("Failed to connect to database: %s", err)
			return
		}
		name := defaultMigrationName
		if len(args) > 0 {
			name = args[0]
		}
		if err := m.Create(name); err != nil {
			fmt.Printf("Failed to create migration: %s", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
