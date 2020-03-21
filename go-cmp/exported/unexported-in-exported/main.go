package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Address struct {
	Province string
	city     string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	u1 := User{"dj", 18, Address{}}
	u2 := User{"dj", 18, Address{}}

	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2, cmpopts.IgnoreUnexported(User{})))
}
