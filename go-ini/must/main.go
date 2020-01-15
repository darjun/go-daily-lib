package main

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	redisPort, err := cfg.Section("redis").Key("port").Int()
	if err != nil {
		fmt.Println("before must, get redis port error:", err)
	} else {
		fmt.Println("before must, get redis port:", redisPort)
	}

	fmt.Println("redis Port:", cfg.Section("redis").Key("port").MustInt(6381))

	redisPort, err = cfg.Section("redis").Key("port").Int()
	if err != nil {
		fmt.Println("after must, get redis port error:", err)
	} else {
		fmt.Println("after must, get redis port:", redisPort)
	}
}