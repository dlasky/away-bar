package main

import (
	"context"
	"dlasky/away-bar/internal"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := &cli.App{
		Name:  "away bar",
		Usage: "a simple GTK status bar",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "css",
				Value: internal.GetConfigPath("style.css"),
				Usage: "specify a css file to use, defaults to .config/awaybar/style.css",
			},
			&cli.StringFlag{
				Name:  "config",
				Value: internal.GetConfigPath("config.hcl"),
				Usage: "specify a config file to use, defaults to .config/awaybar/config.hcl",
			},
		},
		Action: func(c *cli.Context) error {
			App(context.Background())
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
