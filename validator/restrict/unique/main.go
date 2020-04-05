package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v10"
)

type User struct {
	Name    string   `validate:"min=2"`
	Age     int      `validate:"min=18"`
	Hobbies []string `validate:"unique"`
	Friends []User   `validate:"unique=Name"`
}

func main() {
	validate := validator.New()

	f1 := User{
		Name: "dj2",
		Age:  18,
	}
	f2 := User{
		Name: "dj3",
		Age:  18,
	}

	u1 := User{
		Name:    "dj",
		Age:     18,
		Hobbies: []string{"pingpong", "chess", "programming"},
		Friends: []User{f1, f2},
	}
	err := validate.Struct(u1)
	if err != nil {
		fmt.Println(err)
	}

	u2 := User{
		Name:    "dj",
		Age:     18,
		Hobbies: []string{"programming", "programming"},
		Friends: []User{f1, f1},
	}
	err = validate.Struct(u2)
	if err != nil {
		fmt.Println(err)
	}
}
