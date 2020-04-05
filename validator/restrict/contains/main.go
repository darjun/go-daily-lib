package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v10"
)

type User struct {
	Name string `validate:"containsrune=☻"`
	Age  int    `validate:"min=18"`
}

func main() {
	validate := validator.New()

	u1 := User{"d☻j", 18}
	err := validate.Struct(u1)
	if err != nil {
		fmt.Println(err)
	}

	u2 := User{"dj", 18}
	err = validate.Struct(u2)
	if err != nil {
		fmt.Println(err)
	}
}
