package cmd

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/materials-resources/s_prophet/app"
	_ "github.com/materials-resources/s_prophet/internal/catalog"
	_ "github.com/materials-resources/s_prophet/internal/customer"
	_ "github.com/materials-resources/s_prophet/internal/inventory"
	_ "github.com/materials-resources/s_prophet/internal/order"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		cPath := rootCmd.PersistentFlags().Lookup("config").Value

		c, err := app.NewConfigFromPath(cPath.String())

		a, err := app.NewApp(c)
		if err != nil {
			fmt.Println(err)
		}

		// Enable GRPC Reflection for clients
		reflection.Register(a.GetServer())

		a.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
