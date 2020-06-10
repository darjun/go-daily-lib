package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/darjun/go-daily-lib/twirp/get-started/proto"
	"github.com/twitchtv/twirp"
)

func main() {
	client := proto.NewEchoProtobufClient("http://localhost:8080", &http.Client{})

	header := make(http.Header)
	header.Set("Twirp-Token", "test-twirp-token")

	ctx := context.Background()
	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		log.Fatalf("twirp error setting headers: %v", err)
	}

	response, err := client.Say(ctx, &proto.Request{Text: "Hello World"})
	if err != nil {
		log.Fatalf("call say failed: %v", err)
	}
	fmt.Printf("response:%s\n", response.GetText())
}
