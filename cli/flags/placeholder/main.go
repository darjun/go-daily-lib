package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := & cli.App{
		Flags : []cli.Flag {
			&cli.StringFlag{
				Name:"config",
				Usage: "Load configuration from `FILE`",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}