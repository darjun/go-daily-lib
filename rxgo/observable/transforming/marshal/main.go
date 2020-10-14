package main

import (
	"encoding/json"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	observable := rxgo.Just(
		User{
			Name: "dj",
			Age:  18,
		},
		User{
			Name: "jw",
			Age:  20,
		},
	)()

	observable = observable.Marshal(json.Marshal)

	for item := range observable.Observe() {
		fmt.Println(string(item.V.([]byte)))
	}
}
