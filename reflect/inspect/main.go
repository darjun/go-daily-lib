package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type User struct {
	Name    string
	Age     int
	Married bool
}

func (u *User) SetName(n string) {
	u.Name = n
}

func (u *User) SetAge(a int) {
	u.Age = a
}

func inspectStruct(u interface{}) {
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Uint())

		case reflect.Bool:
			fmt.Printf("field:%d type:%s value:%t\n", i, field.Type().Name(), field.Bool())

		case reflect.String:
			fmt.Printf("field:%d type:%s value:%q\n", i, field.Type().Name(), field.String())

		default:
			fmt.Printf("field:%d unhandled kind:%s\n", i, field.Kind())
		}
	}
}

func inspectMap(m interface{}) {
	v := reflect.ValueOf(m)
	for _, k := range v.MapKeys() {
		field := v.MapIndex(k)

		fmt.Printf("%v => %v\n", k.Interface(), field.Interface())
	}
}

func inspectSliceArray(sa interface{}) {
	v := reflect.ValueOf(sa)

	fmt.Printf("%c", '[')
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}
	fmt.Printf("%c\n", ']')
}

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello " + name
}

func inspectFunc(name string, f interface{}) {
	t := reflect.TypeOf(f)
	fmt.Println(name, "input:")
	for i := 0; i < t.NumIn(); i++ {
		t := t.In(i)
		fmt.Print(t.Name())
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Println("output:")
	for i := 0; i < t.NumOut(); i++ {
		t := t.Out(i)
		fmt.Print(t.Name())
		fmt.Print(" ")
	}
	fmt.Println("\n===========")
}

func inspectMethod(o interface{}) {
	t := reflect.TypeOf(o)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		fmt.Println(m)
	}
}

func main() {
	u := User{
		Name:    "dj",
		Age:     18,
		Married: true,
	}

	inspectStruct(u)
	inspectStruct(bytes.Buffer{})

	inspectMap(map[uint32]uint32{
		1: 2,
		3: 4,
	})

	inspectSliceArray([]int{1, 2, 3})
	inspectSliceArray([3]int{4, 5, 6})

	inspectFunc("Add", Add)
	inspectFunc("Greeting", Greeting)

	inspectMethod(&u)
}
