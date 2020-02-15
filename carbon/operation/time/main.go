package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("now is:", now)

	fmt.Println("one second later is:", now.Add(time.Second))
	fmt.Println("one minute later is:", now.Add(time.Minute))
	fmt.Println("one hour later is:", now.Add(time.Hour))

	d, err := time.ParseDuration("3m20s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("3 minutes and 20 seconds later is:", now.Add(d))

	d, err = time.ParseDuration("2h30m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2 hours and 30 minutes later is:", now.Add(d))

	fmt.Println("3 days and 2 hours later is:", now.AddDate(0, 0, 3).Add(time.Hour*2))
}
