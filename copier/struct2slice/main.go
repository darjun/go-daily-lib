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
	user := User{Name: "dj", Age: 18}
	employees := []Employee{}

	copier.Copy(&employees, &user)
	fmt.Printf("%#v\n", employees)
}
