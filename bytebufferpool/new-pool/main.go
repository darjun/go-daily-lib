package main

import (
	"fmt"

	"github.com/valyala/bytebufferpool"
)

func main() {
	joinPool := new(bytebufferpool.Pool)
	b := joinPool.Get()
	b.WriteString("hello")
	b.WriteByte(',')
	b.WriteString(" world!")

	fmt.Println(b.String())

	joinPool.Put(b)
}
