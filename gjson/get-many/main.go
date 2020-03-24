package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `
{
	"name":"dj",
	"age":18,
	"pets": ["cat", "dog"],
	"contact": {
		"phone": "123456789",
		"email": "dj@example.com"
	}
}`

func main() {
	results := gjson.GetMany(json, "name", "age", "pets.#", "contact.phone")
	for _, result := range results {
		fmt.Println(result)
	}
}
