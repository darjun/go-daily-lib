package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	gObj := gabs.New()

	gObj.Set("lee", "info", "name", "first")
	gObj.SetP("darjun", "info.name.last")
	gObj.SetJSONPointer(18, "/info/age")

	fmt.Println(gObj.String())
	// fmt.Println(gObj.StringIndent("", "  "))
}
