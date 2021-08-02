package main

import (
	"fmt"
	"strconv"

	"github.com/MichaelDarr/stack/backend/internal/config"
	"github.com/MichaelDarr/stack/backend/internal/migrate"
	"github.com/spf13/cobra"
)

var forceCmd = &cobra.Command{
	Use:   "force V",
	Short: "Set version V but don't run migration (ignores dirty state)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Failed to read version argument V")
			return
		}
		cfg := config.New()
		m, err := migrate.New(cfg.Postgres.MigrationFilepath, cfg.Postgres.URL())
		if err != nil {
			fmt.Printf("Failed to connect to database: %s", err)
			return
		}
		if err := m.Force(int(version)); err != nil {
			fmt.Println("Failed to force version:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(forceCmd)
}
