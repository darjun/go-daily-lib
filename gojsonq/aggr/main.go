package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("./data.json").From("items")

	fmt.Println("Total Count:", gq.Sum("count"))
	fmt.Println("Min Price:", gq.Min("price"))
	fmt.Println("Max Price:", gq.Max("price"))
	fmt.Println("Avg Price:", gq.Avg("price"))
}
