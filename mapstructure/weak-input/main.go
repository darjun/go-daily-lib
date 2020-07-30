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
		"age":    "18",
		"emails": []int{1, 2, 3},
	}

	var p Person
	err := mapstructure.WeakDecode(m, &p)
	if err == nil {
		fmt.Println("person:", p)
	} else {
		fmt.Println(err.Error())
	}
}
