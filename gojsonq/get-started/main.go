package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	content := `{
	"user": {
		"name": "dj",
		"age": 18,
		"address": {
			"provice": "shanghai",
			"district": "xuhui"
		},
		"hobbies":["chess", "programming", "game"]
	}
}`

	gq := gojsonq.New().FromString(content)
	district := gq.Find("user.address.district")
	fmt.Println(district)

	gq.Reset()

	hobby := gq.Find("user.hobbies.[0]")
	fmt.Println(hobby)
}