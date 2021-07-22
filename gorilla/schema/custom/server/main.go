package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Info</title>
</head>
<body>
<form action="/info" method="post">
	<label>name:</label>
	<input name="name"><br>
	<label>age:</label>
	<input name="age"><br>
    <label>hobbies:</label>
    <input name="hobbies"><br>
	<button type="submit">提交</button>
</form>
</body>
</html>`)
}

type Person struct {
	Name    string   `schema:"name"`
	Age     int      `schema:"age"`
	Hobbies []string `schema:"hobbies"`
}

var (
	decoder = schema.NewDecoder()
)

func init() {
	decoder.RegisterConverter([]string{}, func(s string) reflect.Value {
		return reflect.ValueOf(strings.Split(s, ","))
	})
}

func doinfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := Person{}
	decoder.Decode(&p, r.PostForm)

	fmt.Println(p)
	fmt.Fprintf(w, "Name:%s Age:%d Hobbies:%v", p.Name, p.Age, p.Hobbies)
}

func main() {
	r := mux.NewRouter()
	r.Handle("/info", handlers.MethodHandler{
		"GET":  http.HandlerFunc(info),
		"POST": http.HandlerFunc(doinfo),
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
