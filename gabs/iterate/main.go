package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{
		"user": {
			"name": "dj",
			"age": 18,
			"members": [
				{
					"name": "hjw",
					"age": 20,
					"relation": "spouse"
				},
				{
					"name": "lizi",
					"age": 3,
					"relation": "son"
				}
			]
		}
	}`))

	for k, v := range jObj.S("user").ChildrenMap() {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	fmt.Println()

	for i, v := range jObj.S("user", "members", "*").Children() {
		fmt.Printf("member %d: %v\n", i+1, v)
	}
}
