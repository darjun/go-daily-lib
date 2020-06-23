package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.AppHelpTemplate = `NAME:
	{{.Name}} - {{.Usage}}
USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
 {{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
	{{.Copyright}}
	{{end}}{{if .Version}}
VERSION:
	{{.Version}}
{{end}}
 `

	app := &cli.App{
		Authors: []*cli.Author{
			{
				Name:  "dj",
				Email: "darjun@126.com",
			},
		},
	}
	app.Run(os.Args)
}
