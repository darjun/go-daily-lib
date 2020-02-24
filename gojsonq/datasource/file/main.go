package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("./data.json")

	fmt.Println(gq.Find("items.[1].price"))
}
