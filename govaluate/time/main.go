package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, _ := govaluate.NewEvaluableExpression("'2014-01-02' > '2014-01-01 23:59:59'")
	result, _ := expr.Evaluate(nil)
	fmt.Println(result)
}
