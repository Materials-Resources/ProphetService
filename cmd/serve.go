package cmd

import (
	"github.com/materials-resources/s_prophet/pkg/server"
	"github.com/spf13/cobra"
	"log"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Serve()
		if err != nil {
			log.Fatalf("Could not start server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
