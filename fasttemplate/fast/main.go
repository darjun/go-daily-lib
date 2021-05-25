package main

import (
	"fmt"

	"github.com/valyala/fasttemplate"
)

func main() {
	template := `name: [name]
age: [age]`
	s := fasttemplate.ExecuteString(template, "[", "]", map[string]interface{}{
		"name": "dj",
		"age":  "18",
	})
	fmt.Println(s)
}
