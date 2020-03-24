package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `
{
	"name":{"first":"Tom", "last": "Anderson"},
	"age": 37,
	"children": ["Sara", "Alex", "Jack"],
	"fav.movie": "Dear Hunter",
	"friends": [
		{"first": "Dale", "last":"Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	]
}`

func main() {
	fmt.Println("children in reverse order:", gjson.Get(json, "children|@reverse"))
	fmt.Println("last child:", gjson.Get(json, "children|@reverse|0"))
	fmt.Println("friends with @ugly:", gjson.Get(json, "friends|@ugly"))
	fmt.Println("friends with @pretty:", gjson.Get(json, "friends|@pretty"))
	fmt.Println("root json:", gjson.Get(json, "@this"))

	nestedJSON := `{"nested": ["one", "two", ["three", "four"]]}`
	fmt.Println("@flatten:", gjson.Get(nestedJSON, "nested|@flatten"))

	userJSON := `{"info":[{"name":"dj", "age":18},{"phone":"123456789","email":"dj@example.com"}]}`
	fmt.Println("@join:", gjson.Get(userJSON, "info|@join"))
}
