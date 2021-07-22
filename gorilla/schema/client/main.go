package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

type User struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
}

var (
	encoder = schema.NewEncoder()
)

func main() {
	client := &http.Client{}
	form := url.Values{}

	u := &User{
		Username: "dj",
		Password: "handsome",
	}
	encoder.Encode(u, form)

	res, _ := client.PostForm("http://localhost:8080/login", form)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	res.Body.Close()
}
