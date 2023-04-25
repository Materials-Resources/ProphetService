package cmd

import (
	"github.com/spf13/cobra"
	"materials-resources.com/msv/prophet/pkg/server"
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
