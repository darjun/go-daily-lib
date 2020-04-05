package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v10"
)

func processErr(err error) {
	if err == nil {
		return
	}

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return
	}

	validationErrs := err.(validator.ValidationErrors)
	for _, validationErr := range validationErrs {
		fmt.Println(validationErr)
	}
}

func main() {
	validate := validator.New()

	err := validate.Struct(1)
	processErr(err)

	err = validate.VarWithValue(1, 2, "eqfield")
	processErr(err)
}
