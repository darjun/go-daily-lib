package main

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func main() {
	vancouver, _ := carbon.Today("Asia/Shanghai")
	london, _ := carbon.Today("Asia/Hong_Kong")
	fmt.Println(vancouver.DiffInSeconds(london, true)) // 0

	ottawa, _ := carbon.CreateFromDate(2000, 1, 1, "America/Toronto")
	vancouver, _ = carbon.CreateFromDate(2000, 1, 1, "America/Vancouver")
	fmt.Println(ottawa.DiffInHours(vancouver, true)) // 3

	fmt.Println(ottawa.DiffInHours(vancouver, false)) // 3
	fmt.Println(vancouver.DiffInHours(ottawa, false)) // -3

	t, _ := carbon.CreateFromDate(2012, 1, 31, "UTC")
	fmt.Println(t.DiffInDays(t.AddMonth(), true))  // 31
	fmt.Println(t.DiffInDays(t.SubMonth(), false)) // -31

	t, _ = carbon.CreateFromDate(2012, 4, 30, "UTC")
	fmt.Println(t.DiffInDays(t.AddMonth(), true)) // 30
	fmt.Println(t.DiffInDays(t.AddWeek(), true))  // 7

	t, _ = carbon.CreateFromTime(10, 1, 1, 0, "UTC")
	fmt.Println(t.DiffInMinutes(t.AddSeconds(59), true))  // 0
	fmt.Println(t.DiffInMinutes(t.AddSeconds(60), true))  // 1
	fmt.Println(t.DiffInMinutes(t.AddSeconds(119), true)) // 1
	fmt.Println(t.DiffInMinutes(t.AddSeconds(120), true)) // 2
}
