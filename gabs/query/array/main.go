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

	cnt, _ := jObj.ArrayCount("user", "members")
	fmt.Println("member count:", cnt)
	cnt, _ = jObj.ArrayCount("user", "hobbies")
	fmt.Println("hobby count:", cnt)

	ele, _ := jObj.ArrayElement(0, "user", "members")
	fmt.Println("first member:", ele)
	ele, _ = jObj.ArrayElement(1, "user", "hobbies")
	fmt.Println("second hobby:", ele)
}
