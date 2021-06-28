package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Name   string
	Age    int
	Emails []string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Emails, validation.Each(is.Email)))
}

func main() {
	u := &User{
		Name: "dj",
		Age:  18,
		Emails: []string{
			"darjun@126.com",
			"don't know",
		},
	}
	fmt.Println(validation.Validate(u))
}
