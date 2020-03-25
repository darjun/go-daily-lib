package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

const json = `{"name":{"first":"li","last":"dj"},"age":18}`

func main() {
	value, _ := sjson.Set(json, "name.last", "dajun")
	fmt.Println(value)
}
