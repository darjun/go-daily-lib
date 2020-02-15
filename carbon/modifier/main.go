package main

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	t := carbon.Now()
	fmt.Printf("Start of day:%s\n", t.StartOfDay())
	fmt.Printf("End of day:%s\n", t.EndOfDay())
	fmt.Printf("Start of month:%s\n", t.StartOfMonth())
	fmt.Printf("End of month:%s\n", t.EndOfMonth())
	fmt.Printf("Start of year:%s\n", t.StartOfYear())
	fmt.Printf("End of year:%s\n", t.EndOfYear())
	fmt.Printf("Start of decade:%s\n", t.StartOfDecade())
	fmt.Printf("End of decade:%s\n", t.EndOfDecade())
	fmt.Printf("Start of century:%s\n", t.StartOfCentury())
	fmt.Printf("End of century:%s\n", t.EndOfCentury())
	fmt.Printf("Start of week:%s\n", t.StartOfWeek())
	fmt.Printf("End of week:%s\n", t.EndOfWeek())
	fmt.Printf("Next:%s\n", t.Next(time.Wednesday))
	fmt.Printf("Previous:%s\n", t.Previous(time.Wednesday))
}
