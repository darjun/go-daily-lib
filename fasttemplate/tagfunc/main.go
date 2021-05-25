package main

import (
	"fmt"
	"io"

	"github.com/valyala/fasttemplate"
)

func main() {
	template := `name: {{name}}
age: {{age}}`
	t := fasttemplate.New(template, "{{", "}}")
	s := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "name":
			return w.Write([]byte("dj"))
		case "age":
			return w.Write([]byte("18"))
		default:
			return 0, nil
		}
	})

	fmt.Println(s)
}
