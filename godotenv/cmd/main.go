package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("name"))
	fmt.Println(os.Getenv("version"))
}
