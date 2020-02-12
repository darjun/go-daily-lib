package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("version: ", os.Getenv("version"))
}
