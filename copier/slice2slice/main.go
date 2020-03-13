package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	users := []User{
		{Name: "dj", Age: 18},
		{Name: "dj2", Age: 18},
	}
	employees := []Employee{}

	copier.Copy(&employees, &users)
	fmt.Printf("%#v\n", employees)
}
