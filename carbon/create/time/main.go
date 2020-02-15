package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Japan")
	if err != nil {
		log.Fatal("failed to load location: ", err)
	}

	d := time.Date(2020, time.July, 24, 20, 0, 0, 0, loc)
	fmt.Printf("The opening ceremony of next olympics will start at %s in Japan\n", d)
}
