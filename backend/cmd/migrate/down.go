package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down [n]",
	Short: "Downgrade migrations",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("downgrading all the way")
			return
		}
		fmt.Println("downgrading", args[0], "times")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
