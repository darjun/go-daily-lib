package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  int
	Role string
}

type Employee struct {
	Name      string
	Age       int
	SuperRole string
}

func (e *Employee) Role(role string) {
	e.SuperRole = "Super" + role
}

func main() {
	user := User{Name: "dj", Age: 18, Role: "Admin"}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}
