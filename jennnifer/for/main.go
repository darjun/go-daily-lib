package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")

	f.Func().Id("main").Params().Block(
		Var().Id("sum").Int(),

		For(
			Id("i").Op(":=").Lit(0),
			Id("i").Op("<=").Lit(100),
			Id("i").Op("++"),
		).Block(
			Id("sum").Op("+=").Id("i"),
		),

		Qual("fmt", "Println").Call(Id("sum")),
	)

	fmt.Printf("%#v\n", f)
}
