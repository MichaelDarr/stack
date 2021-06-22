package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up [n]",
	Short: "Upgrade migrations",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("upgrading all the way")
			return
		}
		fmt.Println("upgrading", args[0], "times")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
