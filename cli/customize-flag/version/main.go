package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "version",
		Version: "v1.0.0",
	}
	app.Run(os.Args)
}
