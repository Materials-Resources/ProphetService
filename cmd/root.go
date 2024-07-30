package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "prophet",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCmd.PersistentFlags().StringP(
		"config",
		"c",
		"config.yml",
		"Path to configuration file",
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
