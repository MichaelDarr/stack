package main

import "github.com/spf13/cobra"

// rootCmd is the base command which all others are added to
var rootCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A migration CLI",
}

func main() {
	rootCmd.Execute()
}
