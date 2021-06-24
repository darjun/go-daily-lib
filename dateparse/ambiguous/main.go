package main

import (
	"fmt"
	"log"

	"github.com/araddon/dateparse"
)

func main() {
	t, err := dateparse.ParseStrict("3/1/2014")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
