package main

import (
	"fmt"

	"github.com/valyala/bytebufferpool"
)

func main() {
	b := bytebufferpool.Get()
	b.WriteString("hello")
	b.WriteByte(',')
	b.WriteString(" world!")

	fmt.Println(b.String())

	bytebufferpool.Put(b)
}
