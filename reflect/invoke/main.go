package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello " + name
}

func invoke(f interface{}, args ...interface{}) {
	v := reflect.ValueOf(f)

	argsV := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		argsV = append(argsV, reflect.ValueOf(arg))
	}

	rets := v.Call(argsV)

	fmt.Println("ret:")
	for _, ret := range rets {
		fmt.Println(ret.Interface())
	}
}

type M struct {
	a, b int
	op   rune
}

func (m M) Op() int {
	switch m.op {
	case '+':
		return m.a + m.b

	case '-':
		return m.a - m.b

	case '*':
		return m.a * m.b

	case '/':
		return m.a / m.b

	default:
		panic("invalid op")
	}
}

type Math struct {
	a, b int
}

func (m Math) Add() int {
	return m.a + m.b
}

func (m Math) Sub() int {
	return m.a - m.b
}

func (m Math) Mul() int {
	return m.a * m.b
}

func (m Math) Div() int {
	return m.a / m.b
}

func invokeMethod(obj interface{}, name string, args ...interface{}) {
	v := reflect.ValueOf(obj)
	m := v.MethodByName(name)

	argsV := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		argsV = append(argsV, reflect.ValueOf(arg))
	}

	rets := m.Call(argsV)

	fmt.Println("ret:")
	for _, ret := range rets {
		fmt.Println(ret.Interface())
	}
}

func main() {
	invoke(Add, 1, 2)
	invoke(Greeting, "dj")

	m1 := M{1, 2, '+'}
	m2 := M{3, 4, '-'}
	m3 := M{5, 6, '*'}
	m4 := M{8, 2, '/'}
	invoke(m1.Op)
	invoke(m2.Op)
	invoke(m3.Op)
	invoke(m4.Op)

	m := Math{a: 10, b: 2}
	invokeMethod(m, "Add")
	invokeMethod(m, "Sub")
	invokeMethod(m, "Mul")
	invokeMethod(m, "Div")
}
