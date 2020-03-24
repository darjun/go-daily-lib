package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `
{"name": "Gilbert", "age": 61}
{"name": "Alexa", "age": 34}
{"name": "May", "age": 57}
{"name": "Deloise", "age": 44}`

func main() {
	fmt.Println(gjson.Get(json, "..#"))
	fmt.Println(gjson.Get(json, "..1"))
	fmt.Println(gjson.Get(json, "..#.name"))
	fmt.Println(gjson.Get(json, `..#(name="May").age`))
}
