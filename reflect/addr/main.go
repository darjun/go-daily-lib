package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	x := 2

	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())

	d := c.Elem()
	fmt.Println(d.CanAddr())

	s := []int{1, 2, 3}
	sv := reflect.ValueOf(s)
	e := sv.Index(1)
	fmt.Println(e.CanAddr())

	u := &User{Name: "dj", Age: 18}
	uv := reflect.ValueOf(u)
	f := uv.Elem().Field(0)
	fmt.Println(f.CanAddr())
}
