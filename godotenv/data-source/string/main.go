package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	content := `
	name: awesome web
	version: 0.0.1
	`
	myEnv, err := godotenv.Unmarshal(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", myEnv["name"])
	fmt.Println("version: ", myEnv["version"])
}
