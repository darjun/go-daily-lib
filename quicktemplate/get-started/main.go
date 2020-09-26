package main

import (
	"fmt"

	"github.com/darjun/go-daily-lib/quicktemplate/get-started/templates"
)

func main() {
	fmt.Print(templates.Greeting("dj", 5))
}
