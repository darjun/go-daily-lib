package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, r *http.Request) {
	ptTemplate.ExecuteTemplate(w, "login.tpl", nil)
}

func doLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username != "darjun" || password != "handsome" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	token := fmt.Sprintf("username=%s&password=%s", username, password)
	data := base64.StdEncoding.EncodeToString([]byte(token))
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    data,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func InitLoginRouter(r *mux.Router) {
	ls := r.PathPrefix("/login").Subrouter()
	ls.Methods("GET").HandlerFunc(login)
	ls.Methods("POST").HandlerFunc(doLogin)
}
