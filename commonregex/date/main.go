package main

import (
	"fmt"

	"github.com/mingrammer/commonregex"
)

func main() {
	text := `commonregex support many date formats, like 09.11.2020, Sep 11th 2020 and so on.`
	dateList := commonregex.Date(text)

	fmt.Println(dateList)
}
