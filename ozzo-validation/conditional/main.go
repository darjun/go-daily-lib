package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Name    string
	Age     int
	Student bool
	School  string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 10)),
		validation.Field(&u.Age, validation.Required, validation.Min(1), validation.Max(200)),
		validation.Field(&u.School, validation.When(u.Student, validation.Required, validation.Length(10, 20))))
}

func main() {
	u1 := &User{
		Name:    "dj",
		Age:     18,
		Student: true,
	}

	u2 := &User{
		Name: "lidajun",
		Age:  31,
	}

	fmt.Println("user1:", validation.Validate(u1))
	fmt.Println("user2:", validation.Validate(u2))
}
