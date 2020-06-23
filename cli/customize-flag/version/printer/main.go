package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	Revision = "0cebd6e32a4e7094bbdbf150a1c2ffa56c34e91b"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
	}

	app := &cli.App{
		Name:    "version",
		Version: "v1.0.0",
	}
	app.Run(os.Args)
}
