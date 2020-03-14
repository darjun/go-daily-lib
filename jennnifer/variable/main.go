package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		// Var().Id("greeting").Op("=").Lit("Hello World"),
		Id("greeting").Op(":=").Lit("Hello World"),
		Qual("fmt", "Println").Call(Id("greeting")),
	)
	fmt.Printf("%#v\n", f)
}
