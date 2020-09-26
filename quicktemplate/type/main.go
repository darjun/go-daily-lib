package main

import (
	"fmt"

	"github.com/darjun/go-daily-lib/quicktemplate/type/templates"
)

func main() {
	fmt.Println(templates.Types(1, 5.75, []byte{'a', 'b', 'c'}, "hello"))
}
