package main

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func main() {
	now := carbon.Now()

	fmt.Println("now is:", now)

	fmt.Println("one second later is:", now.AddSecond())
	fmt.Println("one minute later is:", now.AddMinute())
	fmt.Println("one hour later is:", now.AddHour())
	fmt.Println("3 minutes and 20 seconds later is:", now.AddMinutes(3).AddSeconds(20))
	fmt.Println("2 hours and 30 minutes later is:", now.AddHours(2).AddMinutes(30))
	fmt.Println("3 days and 2 hours later is:", now.AddDays(3).AddHours(2))
}
