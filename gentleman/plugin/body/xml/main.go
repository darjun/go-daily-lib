package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

type User struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	cli := gentleman.New()
	cli.URL("http://httpbin.org/post")

	req := cli.Request()
	req.Method("POST")

	u := User{Name: "dj", Age: 18}
	req.Use(body.XML(u))

	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}

	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
