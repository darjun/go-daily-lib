package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func main() {
	name := "darjun"

	err := validation.Validate(name,
		validation.Required,
		validation.Length(2, 10),
		is.URL)
	fmt.Println(err)
}
