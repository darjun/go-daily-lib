package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Email string
	Phone string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.When(u.Phone == "", validation.Required.Error("Either email or phone is required."), is.Email)),
		validation.Field(&u.Phone, validation.When(u.Email == "", validation.Required.Error("Either email or phone is required."), is.Alphanumeric)))
}

func main() {
	u1 := &User{}

	u2 := &User{
		Email: "darjun@126.com",
	}

	u3 := &User{
		Phone: "17301251652",
	}

	u4 := &User{
		Email: "darjun@126.com",
		Phone: "17301251652",
	}

	fmt.Println("user1:", validation.Validate(u1))
	fmt.Println("user2:", validation.Validate(u2))
	fmt.Println("user3:", validation.Validate(u3))
	fmt.Println("user4:", validation.Validate(u4))
}
