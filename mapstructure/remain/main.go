package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name  string
	Age   int
	Job   string
	Other map[string]interface{} `mapstructure:",remain"`
}

func main() {
	data := `
		{ 
			"name": "dj",
			"age":18,
			"job":"programmer",
			"height":"1.8m",
			"handsome": true
		}
	`

	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatal(err)
	}

	var p Person
	mapstructure.Decode(m, &p)
	fmt.Println("other", p.Other)
}
