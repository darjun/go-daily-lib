package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	user := `{"name":"dj","age":18}`
	newValue, err := sjson.Set(user, "na?e", "dajun")
	fmt.Println(err, newValue)
}
