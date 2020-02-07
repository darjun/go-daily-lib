package main

import (
	"log"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "dj",
		Age:  18,
	}

	log.SetPrefix("Login: ")
	log.Printf("%s login, age:%d", u.Name, u.Age)
}
