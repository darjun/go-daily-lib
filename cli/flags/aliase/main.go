package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"language", "l"},
				Value:   "english",
				Usage:   "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "world"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}

			if c.String("lang") == "english" {
				fmt.Println("hello", name)
			} else {
				fmt.Println("你好", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
