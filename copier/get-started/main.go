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
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}
