package main

import (
	"github.com/Jeffail/gabs/v2"
	"fmt"
)

func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{
		"info": {
			"name": {
				"first": "lee",
				"last": "darjun"
			},
			"age": 18,
			"hobbies": [
				"game",
				"programming"
			]
		}
		}`))

	fmt.Println("first name: ", jObj.Search("info", "name", "first").Data().(string))
	fmt.Println("second name: ", jObj.Path("info.name.last").Data().(string))
	gObj, _ := jObj.JSONPointer("/info/age")
	fmt.Println("age: ", gObj.Data().(float64))
	fmt.Println("one hobby: ", jObj.Path("info.hobbies.1").Data().(string))
}