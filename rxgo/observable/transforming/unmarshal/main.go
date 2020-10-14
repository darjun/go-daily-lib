package main

import (
	"context"
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
		`{"name":"dj","age":18}`,
		`{"name":"jw","age":20}`,
	)()

	observable = observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		return []byte(i.(string)), nil
	}).Unmarshal(json.Unmarshal, func() interface{} {
		return &User{}
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
