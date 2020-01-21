package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	p := new(int)
	*p = 8
	fmt.Println(cast.ToInt(p))   // 8

	pp := &p
	fmt.Println(cast.ToInt(pp))  // 8
}