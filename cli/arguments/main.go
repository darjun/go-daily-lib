package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "arguments",
		Usage: "arguments example",
		Action: func(c *cli.Context) error {
			for i := 0; i < c.NArg(); i++ {
				fmt.Printf("%d: %s\n", i+1, c.Args().Get(i))
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
