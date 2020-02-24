package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	file, err := os.OpenFile("./data.json", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	gq := gojsonq.New().Reader(file)

	fmt.Println(gq.Find("items.[1].price"))
}
