package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("./data.json")

	r1 := gq.From("items").Select("id", "name").Offset(0).Limit(3).Get()
	fmt.Println("First Page:", r1)

	gq.Reset()

	r2 := gq.From("items").Select("id", "name").Offset(3).Limit(3).Get()
	fmt.Println("Second Page:", r2)
}
