package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{"user":{"name": "dj","age": 18}}`))

	fmt.Printf("has name? %t\n", jObj.Exists("user", "name"))

	fmt.Printf("has age? %t\n", jObj.ExistsP("user.age"))

	fmt.Printf("has job? %t\n", jObj.Exists("user", "job"))
}
