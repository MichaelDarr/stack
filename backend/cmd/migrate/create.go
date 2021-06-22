package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const defaultMigrationName = "new_migration"

var createCmd = &cobra.Command{
	Use:   "create [migration name]",
	Short: "Create a migration",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		name := defaultMigrationName
		if len(args) > 0 {
			name = args[0]
		}
		fmt.Println("creating migration:", name)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
