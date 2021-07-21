package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Login</title>
</head>
<body>
<form action="/login" method="post">
	<label>Username:</label>
	<input name="username"><br>
	<label>Password:</label>
	<input name="password" type="password"><br>
	<button type="submit">登录</button>
</form>
</body>
</html>`)
}

func dologin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 301)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.Methods("GET").Path("/login").HandlerFunc(login)
	r.Methods("POST").Path("/login").Handler(handlers.ContentTypeHandler(http.HandlerFunc(dologin), "application/x-www-from-urlencoded"))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
