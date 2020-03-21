package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{4, 3, 2, 1}
	fmt.Println("s1 equals s2?", cmp.Equal(s1, s2))
	fmt.Println("s1 equals s2 with option?", cmp.Equal(s1, s2, cmpopts.SortSlices(func(i, j int) bool { return i < j })))

	m1 := map[int]int{1: 10, 2: 20, 3: 30}
	m2 := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println("m1 equals m2?", cmp.Equal(m1, m2))
	fmt.Println("m1 equals m2 with option?", cmp.Equal(m1, m2, cmpopts.SortMaps(func(i, j int) bool { return i < j })))
}
