package main

import (
	"fmt"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

var (
	ptTemplate *template.Template
)

const (
	clientKey    = "dfc0084166d0ef8a9aac"
	clientSecret = "75a7af7446b893707299595d1a4718b7c81174a8"
)

func init() {
	ptTemplate = template.Must(template.New("").ParseGlob("tpls/*.tpl"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ptTemplate.ExecuteTemplate(w, "login.tpl", nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Redirect(w, r, "/login/github", http.StatusTemporaryRedirect)
		return
	}
	ptTemplate.ExecuteTemplate(w, "home.tpl", user)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	ptTemplate.ExecuteTemplate(w, "home.tpl", user)
}

func main() {
	githubProvider := github.New(clientKey, clientSecret, "http://localhost:8080/auth/github/callback")
	goth.UseProviders(githubProvider)
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login/github", LoginHandler)
	r.HandleFunc("/logout/github", LogoutHandler)
	r.HandleFunc("/auth/github", AuthHandler)
	r.HandleFunc("/auth/github/callback", CallbackHandler)

	log.Println("listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
