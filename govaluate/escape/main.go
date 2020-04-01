package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, _ := govaluate.NewEvaluableExpression("[response-time] < 100")
	parameters := make(map[string]interface{})
	parameters["response-time"] = 80
	result, _ := expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression("response\\-time < 100")
	parameters = make(map[string]interface{})
	parameters["response-time"] = 80
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression(`response\-time < 100`)
	parameters = make(map[string]interface{})
	parameters["response-time"] = 80
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)
}
