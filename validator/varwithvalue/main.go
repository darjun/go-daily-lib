package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v10"
)

func main() {
	name1 := "dj"
	name2 := "dj2"

	validate := validator.New()
	fmt.Println(validate.VarWithValue(name1, name2, "eqfield"))

	fmt.Println(validate.VarWithValue(name1, name2, "nefield"))
}