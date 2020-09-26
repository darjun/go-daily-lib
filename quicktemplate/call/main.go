package main

import (
	"fmt"

	"github.com/darjun/go-daily-lib/quicktemplate/call/templates"
)

func main() {
	name2score := make(map[string]int)
	name2score["dj"] = 85
	name2score["lizi"] = 96
	name2score["hjw"] = 52

	fmt.Println(templates.ScoreList(name2score))
}
