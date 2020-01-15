package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	cfg := ini.Empty()

	defaultSection := cfg.Section("")
	defaultSection.NewKey("app_name", "awesome web")
	defaultSection.NewKey("log_level", "DEBUG")

	mysqlSection, err := cfg.NewSection("mysql")
	if err != nil {
		fmt.Println("new mysql section failed:", err)
		return
	}
	mysqlSection.NewKey("ip", "127.0.0.1")
	mysqlSection.NewKey("port", "3306")
	mysqlSection.NewKey("user", "root")
	mysqlSection.NewKey("password", "123456")
	mysqlSection.NewKey("database", "awesome")

	redisSection, err := cfg.NewSection("redis")
	if err != nil {
		fmt.Println("new redis section failed:", err)
		return
	}
	redisSection.NewKey("ip", "127.0.0.1")
	redisSection.NewKey("port", "6381")

	err = cfg.SaveTo("my.ini")
	if err != nil {
		fmt.Println("SaveTo failed: ", err)
	}

	err = cfg.SaveToIndent("my-pretty.ini", "\t")
	if err != nil {
		fmt.Println("SaveToIndent failed: ", err)
	}

	cfg.WriteTo(os.Stdout)
	fmt.Println()
	cfg.WriteToIndent(os.Stdout, "\t")
}