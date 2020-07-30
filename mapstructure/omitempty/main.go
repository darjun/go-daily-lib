package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string `mapstructure:",omitempty"`
}

func main() {
	p := &Person{
		Name: "dj",
		Age:  18,
	}

	var m map[string]interface{}
	mapstructure.Decode(p, &m)

	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}
