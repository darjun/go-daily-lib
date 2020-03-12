package main

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
)

func main() {
	m1 := make(map[string]interface{})
	m1["dbs"] = []uint32{2, 3}

	m2 := make(map[string]interface{})
	m2["dbs"] = []int{1}

	if err := mergo.Map(&m1, &m2, mergo.WithOverride, mergo.WithTypeCheck); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m1)
}
