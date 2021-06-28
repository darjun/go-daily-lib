package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var NameRule = []validation.Rule{
	validation.Required,
	is.Alphanumeric,
	validation.Length(10, 20),
}

func main() {
	name1 := "lidajun12345"
	name2 := "lidajun@!#$%"
	name3 := "short"
	name4 := "looooooooooooooooooong"

	fmt.Println("name1:", validation.Validate(name1, NameRule...))
	fmt.Println("name2:", validation.Validate(name2, NameRule...))
	fmt.Println("name3:", validation.Validate(name3, NameRule...))
	fmt.Println("name4:", validation.Validate(name4, NameRule...))
}
