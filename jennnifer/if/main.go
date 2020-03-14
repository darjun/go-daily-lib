package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")

	f.Func().Id("main").Params().Block(
		Id("score").Op(":=").Lit(70),

		If(Id("score").Op(">=").Lit(90)).Block(
			Qual("fmt", "Println").Call(Lit("优秀")),
		).Else().If(Id("score").Op(">=").Lit(80)).Block(
			Qual("fmt", "Println").Call(Lit("良好")),
		).Else().If(Id("score").Op(">=").Lit(60)).Block(
			Qual("fmt", "Println").Call(Lit("及格")),
		).Else().Block(
			Qual("fmt", "Println").Call(Lit("不及格")),
		),
	)

	fmt.Printf("%#v\n", f)
}