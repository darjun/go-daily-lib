package main

import (
	"fmt"
	"math"

	"github.com/google/go-cmp/cmp"
)

type FloatPair struct {
	X float64
	Y float64
}

func main() {
	p1 := FloatPair{X: math.NaN()}
	p2 := FloatPair{X: math.NaN()}
	fmt.Println("p1 equals p2?", cmp.Equal(p1, p2))

	f1 := 0.1
	f2 := 0.2
	f3 := 0.3
	p3 := FloatPair{X: f1 + f2}
	p4 := FloatPair{X: f3}
	fmt.Println("p3 equals p4?", cmp.Equal(p3, p4))

	p5 := FloatPair{X: 0.1 + 0.2}
	p6 := FloatPair{X: 0.3}
	fmt.Println("p5 equals p6?", cmp.Equal(p5, p6))
}
