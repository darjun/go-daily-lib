package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func (u User) Fullname() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	u := User{FirstName: "li", LastName: "dajun", Age: 18}
	parameters := make(map[string]interface{})
	parameters["u"] = u

	expr, _ := govaluate.NewEvaluableExpression("u.FullName()")
	result, _ := expr.Evaluate(parameters)
	fmt.Println("user", result)

	expr, _ = govaluate.NewEvaluableExpression("u.Age > 18")
	result, _ = expr.Evaluate(parameters)
	fmt.Println("age > 18?", result)
}
