package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v10"
)

type RegisterForm struct {
	Name string `validate:"palindrome"`
	Age  int    `validate:"min=18"`
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func CheckPalindrome(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == reverseString(value)
}

func main() {
	validate := validator.New()
	validate.RegisterValidation("palindrome", CheckPalindrome)

	f1 := RegisterForm{
		Name: "djd",
		Age:  18,
	}
	err := validate.Struct(f1)
	if err != nil {
		fmt.Println(err)
	}

	f2 := RegisterForm{
		Name: "dj",
		Age:  18,
	}
	err = validate.Struct(f2)
	if err != nil {
		fmt.Println(err)
	}
}
