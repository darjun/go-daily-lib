package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("toml")
	tomlConfig := []byte(`
app_name = "awesome web"

# possible values: DEBUG, INFO, WARNING, ERROR, FATAL
log_level = "DEBUG"

[mysql]
ip = "127.0.0.1"
port = 3306
user = "dj"
password = 123456
database = "awesome"

[redis]
ip = "127.0.0.1"
port = 7381
`)
	err := viper.ReadConfig(bytes.NewBuffer(tomlConfig))
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	fmt.Println("redis port: ", viper.GetInt("redis.port"))
}
