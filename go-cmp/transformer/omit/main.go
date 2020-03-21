package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type User struct {
	Name string
	Age  int
}

func omitAge(u User) string {
	return u.Name
}

type User2 struct {
	Name    string
	Age     int
	Email   string
	Address string
}

func omitAge2(u User2) User2 {
	return User2{u.Name, 0, u.Email, u.Address}
}

func main() {
	u1 := User{Name: "dj", Age: 18}
	u2 := User{Name: "dj", Age: 28}

	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2, cmp.Transformer("omitAge", omitAge)))

	u3 := User2{Name: "dj", Age: 18, Email: "dj@example.com"}
	u4 := User2{Name: "dj", Age: 28, Email: "dj@example.com"}

	fmt.Println("u3 equals u4?", cmp.Equal(u3, u4, cmp.Transformer("omitAge", omitAge2)))
}
