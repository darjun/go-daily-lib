package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("./data.json")

	r := gq.From("items").Select("id", "name").
		Where("id", "=", 1).OrWhere("id", "=", 2).Get()
	fmt.Println(r)

	gq.Reset()
	r = gq.From("items").Select("id", "name", "count").
		Where("count", ">", 1).Where("price", "<", 100).Get()
	fmt.Println(r)
}
