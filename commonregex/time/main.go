package main

import (
	"fmt"

	"github.com/mingrammer/commonregex"
)

func main() {
	text := `I wake up at 08:30 (aka 08:30am) in the morning, take a snap at 13:00(aka 01:00pm).`
	timeList := commonregex.Time(text)

	fmt.Println(timeList)
}
