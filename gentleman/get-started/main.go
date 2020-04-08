package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
)

func main() {
	cli := gentleman.New()

	cli.URL("https://dog.ceo")

	req := cli.Request()

	req.Path("/api/breeds/image/random")

	req.SetHeader("Client", "gentleman")

	res, err := req.Send()

	if err != nil {
		fmt.Printf("Request error: %v\n", err)
		return
	}

	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Body: %s", res.String())
}
