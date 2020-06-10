package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/darjun/go-daily-lib/twirp/get-started/proto"
)

func main() {
	client := proto.NewEchoProtobufClient("http://localhost:8080", &http.Client{})

	response, err := client.Say(context.Background(), &proto.Request{Text: "Hello World"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response:%s\n", response.GetText())
}
