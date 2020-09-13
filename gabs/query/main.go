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

	fmt.Println("member names: ", jObj.S("user", "members", "*", "name").Data())
	fmt.Println("member ages: ", jObj.S("user", "members", "*", "age").Data())
	fmt.Println("member relations: ", jObj.S("user", "members", "*", "relation").Data())

	fmt.Println("spouse name: ", jObj.S("user", "members", "0", "name").Data().(string))
}
