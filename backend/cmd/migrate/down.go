package main

import (
	"fmt"
	"strconv"

	"github.com/MichaelDarr/shelf/backend/internal/config"
	"github.com/MichaelDarr/shelf/backend/internal/migrate"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down [N]",
	Short: "Apply all or N down migrations",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		m, err := migrate.New(cfg.Postgres.MigrationFilepath, cfg.Postgres.URL())
		if err != nil {
			fmt.Printf("Failed to connect to database: %s", err)
			return
		}
		if len(args) == 0 {
			if err := m.Down(); err != nil {
				fmt.Println("Failed to migrate down:", err)
			}
			return
		}
		n, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("Failed to read limit argument N")
			return
		}
		if err := m.Steps(int(n) * -1); err != nil {
			fmt.Println("Failed to migrate down:", n, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
