package main

import (
	"fmt"

	"github.com/go-ini/ini"
)

type Config struct {
	AppName 	string `ini:"app_name"`
	LogLevel	string `ini:"log_level"`

	MySQL 		MySQLConfig `ini:"mysql"`
	Redis 		RedisConfig `ini:"redis"`
}

type MySQLConfig struct {
	IP			string `ini:"ip"`
	Port		int `ini:"port"`
	User		string `ini:"user"`
	Password	string `ini:"password"`
	Database	string `ini:"database"`
}

type RedisConfig struct {
	IP		string `ini:"ip"`
	Port	int `ini:"port"`
}

func ConfigToStruct() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Println("load my.ini failed: ", err)
		return
	}

	c := Config{}
	cfg.MapTo(&c)

	fmt.Println(c)
}

func StructToConfig() {
	cfg := ini.Empty()

	c := Config {
		AppName: 	"awesome web",
		LogLevel: 	"DEBUG",
		MySQL: MySQLConfig {
			IP: 	"127.0.0.1",
			Port:	3306,
			User:	"root",
			Password:"123456",
			Database:"awesome",
		},
		Redis: RedisConfig {
			IP:		"127.0.0.1",
			Port:	6381,
		},
	}

	err := ini.ReflectFrom(cfg, &c)
	if err != nil {
		fmt.Println("ReflectFrom failed: ", err)
		return
	}

	err = cfg.SaveTo("my-copy.ini")
	if err != nil {
		fmt.Println("SaveTo failed: ", err)
		return
	}
}

func main() {
	// ConfigToStruct()

	StructToConfig()
}