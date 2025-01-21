package main

import (
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/config"
	_ "github.com/materials-resources/s-prophet/internal/billing"
	_ "github.com/materials-resources/s-prophet/internal/catalog"
	_ "github.com/materials-resources/s-prophet/internal/order"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cliApp := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yaml",
				Usage: "Path to configuration file",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "setup",
				Usage: "Initialize service dependencies",
				Action: func(c *cli.Context) error {
					conf, err := config.NewConfig(c.String("config"))
					if err != nil {
						log.Fatalf("failed to read config")
					}

					a, err := app.NewApp(conf)
					if err != nil {
						os.Exit(1)
					}

					err = a.Setup()
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "serve",
				Usage: "Start the server",
				Action: func(c *cli.Context) error {

					conf, err := config.NewConfig(c.String("config"))
					if err != nil {
						log.Fatalf("failed to read config")
					}

					a, err := app.NewApp(conf)
					if err != nil {
						os.Exit(1)
					}

					err = a.Start()
					if err != nil {
						a.Logger.Fatal().Err(err).Msg("failed to start app")
					}
					//defer a.Stop()
					return nil
				},
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
