package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type NetAddr struct {
	IP   string
	Port int
}

func transformLocalhost(a NetAddr) NetAddr {
	if a.IP == "localhost" {
		return NetAddr{IP: "127.0.0.1", Port: a.Port}
	}

	return a
}

func main() {
	a1 := NetAddr{"127.0.0.1", 5000}
	a2 := NetAddr{"localhost", 5000}

	fmt.Println("a1 equals a2?", cmp.Equal(a1, a2, cmp.Transformer("localhost", transformLocalhost)))
}
