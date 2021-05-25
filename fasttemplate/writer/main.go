package main

import (
	"fmt"
	"io"
	"os"

	"github.com/valyala/fasttemplate"
)

func main() {
	template := `name: {{name}}
age: {{age}}`
	t := fasttemplate.New(template, "{{", "}}")
	t.Execute(os.Stdout, map[string]interface{}{
		"name": "dj",
		"age":  "18",
	})

	fmt.Println()

	t.ExecuteFunc(os.Stdout, func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "name":
			return w.Write([]byte("hjw"))
		case "age":
			return w.Write([]byte("20"))
		}

		return 0, nil
	})
}
