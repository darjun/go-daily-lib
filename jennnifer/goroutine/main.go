package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")
	f.Func().Id("generate").Params().Chan().Int().Block(
		Id("out").Op(":=").Make(Chan().Int()),
		Go().Func().Params().Block(
			For(
				Id("i").Op(":=").Lit(1),
				Id("i").Op("<=").Lit(100),
				Id("i").Op("++"),
			).Block(Id("out").Op("<-").Id("i")),
			Close(Id("out")),
		).Call(),
		Return().Id("out"),
	)

	f.Func().Id("double").Params(Id("in").Op("<-").Chan().Int()).Chan().Int().Block(
		Id("out").Op(":=").Make(Chan().Int()),
		Go().Func().Params().Block(
			For().Id("i").Op(":=").Range().Id("in").Block(Id("out").Op("<-").Id("i").Op("*").Lit(2)),
			Close(Id("out")),
		).Call(),
		Return().Id("out"),
	)

	f.Func().Id("main").Params().Block(
		For(
			Id("i").Op(":=").Range().Id("double").Call(Id("generate").Call()),
		).Block(
			Qual("fmt", "Println").Call(Id("i")),
		),
	)

	fmt.Printf("%#v\n", f)
}
