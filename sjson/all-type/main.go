package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	nilJSON, _ := sjson.Set("", "key", nil)
	fmt.Println(nilJSON)

	boolJSON, _ := sjson.Set("", "key", false)
	fmt.Println(boolJSON)

	intJSON, _ := sjson.Set("", "key", 1)
	fmt.Println(intJSON)

	floatJSON, _ := sjson.Set("", "key", 10.5)
	fmt.Println(floatJSON)

	strJSON, _ := sjson.Set("", "key", "hello")
	fmt.Println(strJSON)

	mapJSON, _ := sjson.Set("", "key", map[string]interface{}{"hello": "world"})
	fmt.Println(mapJSON)

	u := User{Name: "dj", Age: 18}
	structJSON, _ := sjson.Set("", "key", u)
	fmt.Println(structJSON)
}
