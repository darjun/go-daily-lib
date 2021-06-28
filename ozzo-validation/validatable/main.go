package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Name   string
	Age    int
	Gender string
	Email  string
}

func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 10)),
		validation.Field(&u.Age, validation.Required, validation.Min(1), validation.Max(200)),
		validation.Field(&u.Gender, validation.Required, validation.In("male", "female")),
		validation.Field(&u.Email, validation.Required, validation.Length(10, 50), is.Email))

	return err
}

func main() {
	u1 := &User{
		Name:   "darjun",
		Age:    18,
		Gender: "male",
		Email:  "darjun@126.com",
	}
	fmt.Println("user1:", validation.Validate(u1, validation.NotNil))

	u2 := &User{
		Name:  "lidajun12345",
		Age:   201,
		Email: "lidajun's email",
	}
	fmt.Println("user2:", validation.Validate(u2, validation.NotNil))
}
