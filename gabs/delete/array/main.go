package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{"user":{"bugs":["crash","panic"],"hobbies":["game","programming"]}}`))

	jObj.ArrayRemove(0, "user", "bugs")
	jObj.ArrayRemoveP(1, "user.hobbies")
	fmt.Println(jObj.String())
}
