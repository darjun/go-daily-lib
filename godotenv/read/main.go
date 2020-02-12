package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])
}
