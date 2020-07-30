package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	m := map[string]interface{}{
		"name": "dj",
		"age":  18,
		"job":  "programmer",
	}

	var p Person
	var metadata mapstructure.Metadata
	mapstructure.DecodeMetadata(m, &p, &metadata)

	fmt.Printf("keys:%#v unused:%#v\n", metadata.Keys, metadata.Unused)
}
