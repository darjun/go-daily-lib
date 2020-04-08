package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	c "gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugin"
)

func main() {
	cli := gentleman.New()
	cli.URL("https://httpbin.org")

	cli.Use(plugin.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		fmt.Println("request")

		h.Next(ctx)
	}))

	cli.Use(plugin.NewResponsePlugin(func(ctx *c.Context, h c.Handler) {
		fmt.Println("response")

		h.Next(ctx)
	}))

	req := cli.Request()
	req.Path("/headers")
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
