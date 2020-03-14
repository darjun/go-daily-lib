package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")

	f.Func().Id("add").Params(Id("a"), Id("b").Int()).Int().Block(
		Return(Id("a").Op("+").Id("b")),
	)

	f.Func().Id("main").Params().Block(
		Id("a").Op(":=").Lit(1),
		Id("b").Op(":=").Lit(2),
		Qual("fmt", "Println").Call(Id("add").Call(Id("a"), Id("b"))),
	)

	fmt.Printf("%#v\n", f)
}
