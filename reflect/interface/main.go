package main

import "fmt"

type Animal interface {
	Speak()
}

type Cat struct {
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("Bark")
}

func main() {
	var a Animal

	a = Cat{}
	a.Speak()

	a = Dog{}
	a.Speak()
}
