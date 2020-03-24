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
	gjson.ForEachLine(json, func(line gjson.Result) bool {
		fmt.Println("name:", gjson.Get(line.String(), "name"))
		return true
	})
}
