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
}
`

func main() {
	fmt.Println("last name:", gjson.Get(json, "name.last"))
	fmt.Println("age:", gjson.Get(json, "age"))
	fmt.Println("children:", gjson.Get(json, "children"))
	fmt.Println("children count:", gjson.Get(json, "children.#"))
	fmt.Println("second child:", gjson.Get(json, "children.1"))
	fmt.Println("third child*:", gjson.Get(json, "child*.2"))
	fmt.Println("first c?ild:", gjson.Get(json, "c?ildren.0"))
	fmt.Println("fave.moive", gjson.Get(json, `fav.\moive`))
	fmt.Println("first name of friends:", gjson.Get(json, "friends.#.first"))
	fmt.Println("last name of second friend:", gjson.Get(json, "friends.1.last"))

	fmt.Println("first name of first friend whose last name is Murphy:", gjson.Get(json, `friends.#(last="Murphy").first`))
	fmt.Println("first name of friends whose last name is Murphy:", gjson.Get(json, `friends.#(last="Murphy")#.first`))
	fmt.Println("last name of friends whose age > 45:", gjson.Get(json, "friends.#(age>45)#.last"))
	fmt.Println("last name of first friend whose first name satisfy D*:", gjson.Get(json, `friends.#(first%"D*").last`))
	fmt.Println("last name of first friend whose first name not satisfy D*:", gjson.Get(json, `friends.#(first!%"D*").last`))
	fmt.Println("first name of friends whose nets has fb:", gjson.Get(json, `friends.#(nets.#(=="fb"))#.first`))
}
