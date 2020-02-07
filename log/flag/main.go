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

	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	log.Printf("%s login, age:%d", u.Name, u.Age)
}
