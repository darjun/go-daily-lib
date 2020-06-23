package main

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Println("Simple help!")
	}

	(&cli.App{}).Run(os.Args)
}
