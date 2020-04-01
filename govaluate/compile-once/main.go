package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, _ := govaluate.NewEvaluableExpression("a + b")
	parameters := make(map[string]interface{})
	parameters["a"] = 1
	parameters["b"] = 2
	result, _ := expr.Evaluate(parameters)
	fmt.Println(result)

	parameters = make(map[string]interface{})
	parameters["a"] = 10
	parameters["b"] = 20
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)
}
