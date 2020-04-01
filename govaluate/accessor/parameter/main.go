package main

import (
	"errors"
	"fmt"

	"github.com/Knetic/govaluate"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func (u User) Get(name string) (interface{}, error) {
	if name == "FullName" {
		return u.FirstName + " " + u.LastName, nil
	}

	return nil, errors.New("unsupported field " + name)
}

func main() {
	u := User{FirstName: "li", LastName: "dajun", Age: 18}
	expr, _ := govaluate.NewEvaluableExpression("FullName")
	result, _ := expr.Eval(u)
	fmt.Println("user", result)
}
