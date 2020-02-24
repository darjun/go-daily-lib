package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("./data.json")

	fmt.Println(gq.From("items").GroupBy("price").Get())

	gq.Reset()

	fmt.Println(gq.From("items").SortBy("price", "desc").Get())
}
