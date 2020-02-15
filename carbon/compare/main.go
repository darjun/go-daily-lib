package main

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func main() {
	t1, _ := carbon.CreateFromDate(2010, 10, 1, "Asia/Shanghai")
	t2, _ := carbon.CreateFromDate(2011, 10, 20, "Asia/Shanghai")

	fmt.Printf("t1 equal to t2: %t\n", t1.Eq(t2))
	fmt.Printf("t1 not equal to t2: %t\n", t1.Ne(t2))

	fmt.Printf("t1 greater than t2: %t\n", t1.Gt(t2))
	fmt.Printf("t1 less than t2: %t\n", t1.Lt(t2))

	t3, _ := carbon.CreateFromDate(2011, 1, 20, "Asia/Shanghai")
	fmt.Printf("t3 between t1 and t2: %t\n", t3.Between(t1, t2, true))

	now := carbon.Now()
	fmt.Printf("Weekday? %t\n", now.IsWeekday())
	fmt.Printf("Weekend? %t\n", now.IsWeekend())
	fmt.Printf("LeapYear? %t\n", now.IsLeapYear())
	fmt.Printf("Past? %t\n", now.IsPast())
	fmt.Printf("Future? %t\n", now.IsFuture())
}
