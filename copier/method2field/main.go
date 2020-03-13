package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  int
}

func (u *User) DoubleAge() int {
	return 2 * u.Age
}

type Employee struct {
	Name      string
	DoubleAge int
	Role      string
}

func main() {
	user := User{Name: "dj", Age: 18}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}
