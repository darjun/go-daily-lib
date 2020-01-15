package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Home dir:", u.HomeDir)
}