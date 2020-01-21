package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	m1 := map[string]string {
		"name": "darjun",
		"job": "developer",
	}

	m2 := map[string]interface{} {
		"name": "jingwen",
		"age": 18,
	}

	m3 := map[interface{}]string {
		"name": "pipi",
		"job": "designer",
	}

	m4 := map[interface{}]interface{} {
		"name": "did",
		"age": 29,
	}

	jsonStr := `{"name":"bibi", "job":"manager"}`

	fmt.Println(cast.ToStringMapString(m1))      // map[job:developer name:darjun]
	fmt.Println(cast.ToStringMapString(m2))      // map[age:18 name:jingwen] 
	fmt.Println(cast.ToStringMapString(m3))      // map[job:designer name:pipi]
	fmt.Println(cast.ToStringMapString(m4))      // map[job:designer name:pipi]
	fmt.Println(cast.ToStringMapString(jsonStr)) // map[job:manager name:bibi]
}