package main

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
)

type redisConfig struct {
	Address string
	Port    int
	DBs     []int
}

var defaultConfig = redisConfig{
	Address: "127.0.0.1",
	Port:    6381,
	DBs:     []int{1},
}

func main() {
	var config redisConfig
	config.DBs = []int{2, 3}

	if err := mergo.Merge(&config, defaultConfig, mergo.WithAppendSlice); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis dbs: ", config.DBs)
}
