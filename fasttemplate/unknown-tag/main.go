package main

import (
	"fmt"

	"github.com/valyala/fasttemplate"
)

func main() {
	template := `name: {{name}}
age: {{age}}`
	t := fasttemplate.New(template, "{{", "}}")
	m := map[string]interface{}{"name": "dj"}
	s1 := t.ExecuteString(m)
	fmt.Println(s1)

	s2 := t.ExecuteStringStd(m)
	fmt.Println(s2)
}
