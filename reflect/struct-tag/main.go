package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := &User{Name: "dj", Age: 18}
	t := reflect.TypeOf(u).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Tag, f.Tag.Get("json"))
	}
}
