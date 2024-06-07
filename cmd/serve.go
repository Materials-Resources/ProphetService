package cmd

import (
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/materials-resources/s-prophet/app"
	_ "github.com/materials-resources/s-prophet/internal/billing"
	_ "github.com/materials-resources/s-prophet/internal/catalog"
	_ "github.com/materials-resources/s-prophet/internal/customer"
	_ "github.com/materials-resources/s-prophet/internal/order"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		cPath := rootCmd.PersistentFlags().Lookup("config").Value

		c := app.NewConfig(cPath.String())

		a, err := app.NewApp(c)
		if err != nil {
			log.Fatalf("failed to intialize app: %v", err)
		}

		a.Start()
		defer a.Stop()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
