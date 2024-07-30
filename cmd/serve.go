package main

import (
	"github.com/materials-resources/s-prophet/config"
	"log"
	"os"

	"github.com/materials-resources/s-prophet/app"
	_ "github.com/materials-resources/s-prophet/internal/billing"
	_ "github.com/materials-resources/s-prophet/internal/catalog"
	_ "github.com/materials-resources/s-prophet/internal/customer"
	_ "github.com/materials-resources/s-prophet/internal/order"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		cPath := rootCmd.PersistentFlags().Lookup("config").Value

		c, err := config.NewConfig(cPath.String())
		if err != nil {
			log.Fatalf("failed to read config")
		}

		a, err := app.NewApp(c)
		if err != nil {
			os.Exit(1)
		}

		err = a.Start()
		if err != nil {
			a.Logger.Fatal().Err(err).Msg("failed to start app")
		}
		defer a.Stop()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
