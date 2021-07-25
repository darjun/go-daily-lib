package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var (
	store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
)

func set(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Values["name"] = "dj"
	session.Values["age"] = 18
	err := store.Save(r, w, session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Hello World")
}

func read(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	fmt.Fprintf(w, "name:%s age:%d\n", session.Values["name"], session.Values["age"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/set", set)
	r.HandleFunc("/read", read)
	log.Fatal(http.ListenAndServe(":8080", r))
}
