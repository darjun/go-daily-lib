package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func validateUser(u *User) error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 10)),
		validation.Field(&u.Age, validation.Required, validation.Min(1), validation.Max(200)),
		validation.Field(&u.Email, validation.Required, validation.Length(10, 50), is.Email))

	return err
}

func main() {
	u1 := &User{
		Name:  "darjun",
		Age:   18,
		Email: "darjun@126.com",
	}
	fmt.Println("user1:", validateUser(u1))

	u2 := &User{
		Name:  "lidajun12345",
		Age:   201,
		Email: "lidajun's email",
	}
	fmt.Println("user2:", validateUser(u2))
}
