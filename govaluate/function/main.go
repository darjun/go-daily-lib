package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	functions := map[string]govaluate.ExpressionFunction{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return length, nil
		},
	}

	exprString := "strlen('teststring')"
	expr, _ := govaluate.NewEvaluableExpressionWithFunctions(exprString, functions)
	result, _ := expr.Evaluate(nil)
	fmt.Println(result)
}
