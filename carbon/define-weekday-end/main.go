package main

import (
	"fmt"
	"log"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	t, err := carbon.Create(2020, 02, 11, 0, 0, 0, 0, "Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}

	t.SetWeekStartsAt(time.Sunday)
	t.SetWeekEndsAt(time.Saturday)
	t.SetWeekendDays([]time.Weekday{time.Monday, time.Tuesday, time.Wednesday})

	fmt.Printf("Today is %s, weekend? %t\n", t.Weekday(), t.IsWeekend())
}
