package main

import (
	"fmt"
	"strconv"

	"github.com/MichaelDarr/stack/backend/internal/config"
	"github.com/MichaelDarr/stack/backend/internal/migrate"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up [N]",
	Short: "Apply all or N up migrations",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		m, err := migrate.New(cfg.Postgres.MigrationFilepath, cfg.Postgres.URL())
		if err != nil {
			fmt.Printf("Failed to connect to database: %s", err)
			return
		}
		if len(args) == 0 {
			if err := m.Up(); err != nil {
				fmt.Println("Failed to migrate up:", err)
			}
			return
		}
		n, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("Failed to read limit argument N")
			return
		}
		if err := m.Steps(int(n)); err != nil {
			fmt.Println("Failed to migrate up:", n, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
