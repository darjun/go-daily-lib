package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")

	f.Type().Id("User").Struct(
		Id("Name").String(),
		Id("Age").Int(),
	)

	f.Func().Params(Id("u").Id("*User")).Id("Greeting").Params().Block(
		Qual("fmt", "Printf").Call(Lit("Hello %s"), Id("u").Dot("Name")),
	)

	f.Func().Id("main").Params().Block(
		Id("u").Op(":=").Id("User").Values(
			Id("Name").Op(":").Lit("dj"),
			Id("Age").Op(":").Lit(18),
		),
		Id("u").Dot("Greeting").Call(),
	)

	fmt.Printf("%#v\n", f)
}
