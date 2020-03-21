package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type Contact struct {
	Phone string
	Email string
}

type User struct {
	Name    string
	Age     int
	Contact *Contact
}

func main() {
	u1 := User{Name: "dj", Age: 18}
	u2 := User{Name: "dj", Age: 18}

	fmt.Println("u1 == u2?", u1 == u2)
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))

	c1 := &Contact{Phone: "123456789", Email: "dj@example.com"}
	c2 := &Contact{Phone: "123456789", Email: "dj@example.com"}

	u1.Contact = c1
	u2.Contact = c1
	fmt.Println("u1 == u2 with same pointer?", u1 == u2)
	fmt.Println("u1 equals u2 with same pointer?", cmp.Equal(u1, u2))

	u2.Contact = c2
	fmt.Println("u1 == u2 with different pointer?", u1 == u2)
	fmt.Println("u1 equals u2 with different pointer?", cmp.Equal(u1, u2))
}
