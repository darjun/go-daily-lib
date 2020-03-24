package main

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}

		if arg == "lower" {
			return strings.ToLower(json)
		}

		return json
	})

	const json = `{"children": ["Sara", "Alex", "Jack"]}`
	fmt.Println(gjson.Get(json, "children|@case:upper"))
	fmt.Println(gjson.Get(json, "children|@case:lower"))
}
