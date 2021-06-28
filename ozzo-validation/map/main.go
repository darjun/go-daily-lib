package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func validateUser(u map[string]interface{}) error {
	err := validation.Validate(u, validation.Map(
		validation.Key("name", validation.Required, validation.Length(2, 10)),
		validation.Key("age", validation.Required, validation.Min(1), validation.Max(200)),
		validation.Key("email", validation.Required, validation.Length(10, 50), is.Email),
	))

	return err
}

func main() {
	u1 := map[string]interface{}{
		"name":  "darjun",
		"age":   18,
		"email": "darjun@126.com",
	}
	fmt.Println("user1:", validateUser(u1))

	u2 := map[string]interface{}{
		"name":  "lidajun12345",
		"age":   201,
		"email": "lidajun's email",
	}
	fmt.Println("user2:", validateUser(u2))
}
