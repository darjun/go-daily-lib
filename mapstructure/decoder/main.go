package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	m := map[string]interface{}{
		"name": 123,
		"age":  "18",
		"job":  "programmer",
	}

	var p Person
	var metadata mapstructure.Metadata

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &p,
		Metadata:         &metadata,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = decoder.Decode(m)
	if err == nil {
		fmt.Println("person:", p)
		fmt.Printf("keys:%#v, unused:%#v\n", metadata.Keys, metadata.Unused)
	} else {
		fmt.Println(err.Error())
	}
}
