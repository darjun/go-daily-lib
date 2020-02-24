package main

import (
	"encoding/json"
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	r := gojsonq.New().File("./data.json").From("items").Select("id", "name").Get()
	data, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(data))
}
