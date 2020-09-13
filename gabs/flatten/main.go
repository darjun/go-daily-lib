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
			],
			"hobbies": ["game", "programming"]
		}
	}`))

	obj, _ := jObj.Flatten()
	fmt.Println(obj)
}
