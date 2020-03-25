package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	var newValue string
	user := `{"name":{"first":"li","last":"dj"},"age":18}`

	newValue, _ = sjson.Delete(user, "name.first")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(user, "name.full")
	fmt.Println(newValue)

	fruits := `{"fruits":["apple", "orange", "banana"]}`

	newValue, _ = sjson.Delete(fruits, "fruits.1")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(fruits, "fruits.-1")
	fmt.Println(newValue)

	newValue, _ = sjson.Delete(fruits, "fruits.5")
	fmt.Println(newValue)
}
