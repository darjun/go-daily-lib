package main

import (
	"fmt"

	"go.uber.org/dig"
)

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) func() *User {
	return func() *User {
		return &User{name, age}
	}
}

type UserParams struct {
	dig.In

	User1 *User `name:"dj"`
	User2 *User `name:"dj2"`
}

func PrintInfo(params UserParams) error {
	fmt.Println("User 1 ===========")
	fmt.Println("Name:", params.User1.Name)
	fmt.Println("Age:", params.User1.Age)

	fmt.Println("User 2 ===========")
	fmt.Println("Name:", params.User2.Name)
	fmt.Println("Age:", params.User2.Age)
	return nil
}

func main() {
	container := dig.New()

	container.Provide(NewUser("dj", 18), dig.Name("dj"))
	container.Provide(NewUser("dj2", 18), dig.Name("dj2"))

	container.Invoke(PrintInfo)
}
