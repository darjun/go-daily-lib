package main

import (
	"fmt"
)

type Animal interface {
	Speak()
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

func main() {
	var a Animal

	a = Cat{Name: "kitty"}
	a.Speak()

	c := a.(Cat)
	fmt.Println(c.Name)
}
