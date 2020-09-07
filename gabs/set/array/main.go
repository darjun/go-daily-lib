package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	gObj := gabs.New()

	arrObj1, _ := gObj.Array("user", "hobbies")
	fmt.Println(arrObj1.String())

	arrObj2, _ := gObj.ArrayP("user.bugs")
	fmt.Println(arrObj2.String())

	gObj.ArrayAppend("game", "user", "hobbies")
	gObj.ArrayAppend("programming", "user", "hobbies")

	gObj.ArrayAppendP("crash", "user.bugs")
	gObj.ArrayAppendP("panic", "user.bugs")
	fmt.Println(gObj.String())
}
