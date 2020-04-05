package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v10"
)

type RegisterForm struct {
	Name      string `validate:"min=2"`
	Age       int    `validate:"min=18"`
	Password  string `validate:"min=10"`
	Password2 string `validate:"eqfield=Password"`
}

func main() {
	validate := validator.New()

	f1 := RegisterForm{
		Name:      "dj",
		Age:       18,
		Password:  "1234567890",
		Password2: "1234567890",
	}
	err := validate.Struct(f1)
	if err != nil {
		fmt.Println(err)
	}

	f2 := RegisterForm{
		Name:      "dj",
		Age:       18,
		Password:  "1234567890",
		Password2: "123",
	}
	err = validate.Struct(f2)
	if err != nil {
		fmt.Println(err)
	}
}
