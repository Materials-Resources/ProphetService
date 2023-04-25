package cmd

import (
	"github.com/materials-resources/ProphetService/pkg/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
