package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Var().Id("a").Op("=").Lit(10),
		Var().Id("b").Op("=").Lit(2),
		Qual("fmt", "Printf").Call(Lit("%d + %d = %d\n"), Id("a"), Id("b"), Id("a").Op("+").Id("b")),
		Qual("fmt", "Printf").Call(Lit("%d + %d = %d\n"), Id("a"), Id("b"), Id("a").Op("-").Id("b")),
		Qual("fmt", "Printf").Call(Lit("%d + %d = %d\n"), Id("a"), Id("b"), Id("a").Op("*").Id("b")),
		Qual("fmt", "Printf").Call(Lit("%d + %d = %d\n"), Id("a"), Id("b"), Id("a").Op("/").Id("b")),
	)
	fmt.Printf("%#v\n", f)
}
