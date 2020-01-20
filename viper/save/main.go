package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.Set("app_name", "awesome web")
	viper.Set("log_level", "DEBUG")
	viper.Set("mysql.ip", "127.0.0.1")
	viper.Set("mysql.port", 3306)
	viper.Set("mysql.user", "root")
	viper.Set("mysql.password", "123456")
	viper.Set("mysql.database", "awesome")

	viper.Set("redis.ip", "127.0.0.1")
	viper.Set("redis.port", 6381)

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}
}
