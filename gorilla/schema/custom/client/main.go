package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

type Person struct {
	Name    string `schema:"name"`
	Age     int    `schema:"age"`
	Hobbies string `schema:"hobbies"`
}

var (
	encoder = schema.NewEncoder()
)

func main() {
	client := &http.Client{}
	form := url.Values{}

	p := &Person{
		Name:    "dj",
		Age:     18,
		Hobbies: "Game,Programming",
	}
	encoder.Encode(p, form)

	res, _ := client.PostForm("http://localhost:8080/info", form)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	res.Body.Close()
}
