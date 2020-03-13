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

func (u *User) DoubleAge() int {
	return u.Age * 2
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
	var (
		user  = User{Name: "dj", Age: 18}
		users = []User{
			{Name: "dj", Age: 18, Role: "Admin"},
			{Name: "dj2", Age: 18, Role: "Dev"},
		}
		employee  = Employee{}
		employees = []Employee{}
	)

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)

	copier.Copy(&employees, &user)
	fmt.Printf("%#v\n", employees)

	// employees = []Employee{}

	copier.Copy(&employees, &users)
	fmt.Printf("%#v\n", employees)
}
