package main

import (
	"fmt"
	"log"
	
	"github.com/mitchellh/go-homedir"
)

func main() {
	homedir.DisableCache = false

	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Home dir:", dir)
}