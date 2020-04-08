package main

import (
  "fmt"
  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman.v2/plugins/headers"
)

func main() {
  cli := gentleman.New()
  cli.URL("https://api.thecatapi.com")

  cli.Use(headers.Set("x-api-key", "479ce48d-db30-46a4-b1a0-91ac4c1477b8"))
  cli.Use(headers.Del("User-Agent"))

  req := cli.Request()
  req.Path("/v1/breeds")
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