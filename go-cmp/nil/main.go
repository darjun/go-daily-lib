package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func main() {
	var s1 []int
	var s2 = make([]int, 0)

	var m1 map[int]int
	var m2 = make(map[int]int)

	fmt.Println("s1 equals s2?", cmp.Equal(s1, s2))
	fmt.Println("m1 equals m2?", cmp.Equal(m1, m2))

	fmt.Println("s1 equals s2 with option?", cmp.Equal(s1, s2, cmpopts.EquateEmpty()))
	fmt.Println("m1 equals m2 with option?", cmp.Equal(m1, m2, cmpopts.EquateEmpty()))
}
