package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darjun/go-daily-lib/twirp/get-started/proto"
)

func main() {
	client := proto.NewEchoProtobufClient("http://localhost:8080", &http.Client{})

	response, _ := client.Say(context.Background(), &proto.Request{Text: "Hello World"})
	fmt.Println("echo:", response.GetText())

	httpResp, _ := http.Get("http://localhost:8080/greeting")
	data, _ := ioutil.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	fmt.Println("greeting:", string(data))

	httpResp, _ = http.Get("http://localhost:8080/greeting?name=dj")
	data, _ = ioutil.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	fmt.Println("greeting:", string(data))
}
