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
	pets := gjson.Get(json, "pets")
	pets.ForEach(func(_, pet gjson.Result) bool {
		fmt.Println(pet)
		return true
	})

	contact := gjson.Get(json, "contact")
	contact.ForEach(func(key, value gjson.Result) bool {
		fmt.Println(key, value)
		return true
	})
}
