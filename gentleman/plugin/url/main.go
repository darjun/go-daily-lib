package main

import (
	"fmt"
	"os"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
)

func main() {
	cli := gentleman.New()
	cli.URL("https://api.thecatapi.com/")

	cli.Use(headers.Set("x-api-key", "479ce48d-db30-46a4-b1a0-91ac4c1477b8"))
	cli.Use(url.Path("/v1/:type"))

	for _, arg := range os.Args[1:] {
		req := cli.Request()
		req.Use(url.Param("type", arg))
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
		fmt.Printf("Body: %s\n", res.String())
	}
}
