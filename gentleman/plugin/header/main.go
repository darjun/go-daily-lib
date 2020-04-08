package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

func main() {
	cli := gentleman.New()
	cli.URL("http://httpbin.org/post")

	data := map[string]string{"foo": "bar"}
	// cli.Use(body.JSON(data))

	req := cli.Request()
	req.Method("POST")
	req.Use(body.JSON(data))

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
