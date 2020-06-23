package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(`%s

WEBSITE: http://darjun.github.io

WECHAT: GoUpUp`, cli.AppHelpTemplate)

	(&cli.App{}).Run(os.Args)
}
