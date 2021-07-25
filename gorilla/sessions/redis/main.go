package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/boj/redistore.v1"
	"html/template"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string
	Count    int
}

var (
	port       = flag.Int("port", 8080, "port to listen")
	ptTemplate *template.Template
	store, _   = redistore.NewRediStore(10, "tcp", ":6379", "", []byte("redis-key"))
)

func GetSessionUser(r *http.Request) *User {
	session, _ := store.Get(r, "user")
	s, ok := session.Values["user"]
	if !ok {
		return nil
	}
	u := &User{}
	json.Unmarshal([]byte(s.(string)), u)
	return u
}

func SaveSessionUser(w http.ResponseWriter, r *http.Request, u *User) {
	session, _ := store.Get(r, "user")
	data, _ := json.Marshal(u)
	session.Values["user"] = string(data)
	store.Save(r, w, session)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	u := GetSessionUser(r)
	ptTemplate.ExecuteTemplate(w, "home.tpl", u)
}

func Login(w http.ResponseWriter, r *http.Request) {
	ptTemplate.ExecuteTemplate(w, "login.tpl", nil)
}

func DoLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username != "darjun" || password != "handsome" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	SaveSessionUser(w, r, &User{Username: username})
	http.Redirect(w, r, "/", http.StatusFound)
}

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	u := GetSessionUser(r)
	if u == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	u.Count++
	SaveSessionUser(w, r, u)
	ptTemplate.ExecuteTemplate(w, "secret.tpl", u)
}

func Logger(h http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, h)
}

func main() {
	r := mux.NewRouter()
	r.Use(Logger)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/secret", SecretHandler)
	r.Handle("/login", handlers.MethodHandler{
		"GET":  http.HandlerFunc(Login),
		"POST": http.HandlerFunc(DoLogin),
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func init() {
	flag.Parse()
	ptTemplate = template.Must(template.New("").ParseGlob("./tpls/*.tpl"))
}
