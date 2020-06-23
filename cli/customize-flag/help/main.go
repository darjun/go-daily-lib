package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "haaaaalp",
		Aliases: []string{"halp"},
		Usage:   "HALP",
	}

	(&cli.App{}).Run(os.Args)
}
