package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "noop",
				Usage: "Usage for noop",
			},
			{
				Name:     "add",
				Category: "template",
				Usage:    "Usage for add",
			},
			{
				Name:     "remove",
				Category: "template",
				Usage:    "Usage for remove",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
