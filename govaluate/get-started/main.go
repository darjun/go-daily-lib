package main

import (
	"fmt"
	"log"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		log.Fatal("syntax error:", err)
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		log.Fatal("evaluate error:", err)
	}

	fmt.Println(result)
}
