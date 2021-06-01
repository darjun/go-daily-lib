package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	age  int
}

func main() {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x) // 3

	d.Set(reflect.ValueOf(4))
	fmt.Println(x) // 4

	d.SetInt(5)
	fmt.Println(x) // 5

	u := &User{Name: "dj", age: 18}
	uv := reflect.ValueOf(u)
	name := uv.Elem().Field(0)
	fmt.Println(name.CanAddr(), name.CanSet())
	age := uv.Elem().Field(1)
	fmt.Println(age.CanAddr(), age.CanSet())

	name.SetString("lidajun")
	fmt.Println(u)
	age.SetInt(20) // 报错
}
