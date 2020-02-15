package main

import (
	"fmt"
	"log"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	c, err := carbon.Create(2020, time.July, 24, 20, 0, 0, 0, "Japan")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The opening ceremony of next olympics will start at %s in Japan\n", c)
}
