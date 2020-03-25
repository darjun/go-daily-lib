package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	fruits := `{"fruits":["apple", "orange", "banana"]}`

	var newValue string
	newValue, _ = sjson.Set(fruits, "fruits.1", "grape")
	fmt.Println(newValue)

	newValue, _ = sjson.Set(fruits, "fruits.3", "pear")
	fmt.Println(newValue)

	newValue, _ = sjson.Set(fruits, "fruits.-1", "strawberry")
	fmt.Println(newValue)

	newValue, _ = sjson.Set(fruits, "fruits.5", "watermelon")
	fmt.Println(newValue)
}
