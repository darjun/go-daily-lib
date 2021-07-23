package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

type User struct {
	Name string
	Age  int
}

var (
	hashKey  = securecookie.GenerateRandomKey(16)
	blockKey = securecookie.GenerateRandomKey(16)
	s        = securecookie.New(hashKey, blockKey)
)

type XMLEncoder struct{}

func (x XMLEncoder) Serialize(src interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := xml.NewEncoder(buf)
	if err := encoder.Encode(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (x XMLEncoder) Deserialize(src []byte, dst interface{}) error {
	dec := xml.NewDecoder(bytes.NewBuffer(src))
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}

func init() {
	s.SetSerializer(XMLEncoder{})
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Name: "dj",
		Age:  18,
	}
	if encoded, err := s.Encode("user", u); err == nil {
		cookie := &http.Cookie{
			Name:     "user",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Fprintln(w, "Hello World")
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("user"); err == nil {
		u := &User{}
		if err = s.Decode("user", cookie.Value, u); err == nil {
			fmt.Fprintf(w, "name:%s age:%d", u.Name, u.Age)
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/set_cookie", SetCookieHandler)
	r.HandleFunc("/read_cookie", ReadCookieHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
