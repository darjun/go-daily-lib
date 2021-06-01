package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
	Name string
}

func main() {
	var f float64 = 3.5
	t1 := reflect.TypeOf(f)
	fmt.Println(t1.String())

	c := Cat{Name: "kitty"}
	t2 := reflect.TypeOf(c)
	fmt.Println(t2.String())

	v1 := reflect.ValueOf(f)
	fmt.Println(v1)
	fmt.Println(v1.String())

	v2 := reflect.ValueOf(c)
	fmt.Println(v2)
	fmt.Println(v2.String())

	fmt.Printf("%T\n", 3) // int
}
