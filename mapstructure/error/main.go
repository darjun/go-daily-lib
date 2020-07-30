package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	m := map[string]interface{}{
		"name":   123,
		"age":    "bad value",
		"emails": []int{1, 2, 3},
	}

	var p Person
	err := mapstructure.Decode(m, &p)
	if err != nil {
		fmt.Println(err.Error())
	}
}
