package main

import (
	"log"
	"os"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)

	_, err := os.Stat("./generated")
	if os.IsNotExist(err) {
		os.Mkdir("./generated", 0666)
	}

	err = f.Save("./generated/main.go")
	if err != nil {
		log.Fatal(err)
	}
}
