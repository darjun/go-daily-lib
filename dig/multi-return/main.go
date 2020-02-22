package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"go.uber.org/dig"
	"gopkg.in/ini.v1"
)

type Option struct {
	ConfigFile string `short:"c" long:"config" description:"Name of config file."`
}

func InitOptionAndConfig() (*Option, *ini.File, error) {
	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		return nil, nil, err
	}

	cfg, err := ini.Load(opt.ConfigFile)
	if err != nil {
		return &opt, nil, err
	}

	return &opt, cfg, err
}

func PrintInfo(cfg *ini.File) {
	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())
}

func main() {
	container := dig.New()

	container.Provide(InitOptionAndConfig)
	container.Invoke(PrintInfo)
}
