package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

type User struct {
	Name string
	Age  int
}

var (
	prevCookie    unsafe.Pointer
	currentCookie unsafe.Pointer
)

func init() {
	prevCookie = unsafe.Pointer(securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	))
	currentCookie = unsafe.Pointer(securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	))
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Name: "dj",
		Age:  18,
	}
	if encoded, err := securecookie.EncodeMulti(
		"user", u,
		(*securecookie.SecureCookie)(atomic.LoadPointer(&currentCookie)),
	); err == nil {
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
		if err = securecookie.DecodeMulti(
			"user", cookie.Value, u,
			(*securecookie.SecureCookie)(atomic.LoadPointer(&currentCookie)),
			(*securecookie.SecureCookie)(atomic.LoadPointer(&prevCookie)),
		); err == nil {
			fmt.Fprintf(w, "name:%s age:%d", u.Name, u.Age)
		} else {
			fmt.Fprintf(w, "read cookie error:%v", err)
		}
	}
}

func main() {
	r := mux.NewRouter()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go RotateKey(ctx)

	r.HandleFunc("/set_cookie", SetCookieHandler)
	r.HandleFunc("/read_cookie", ReadCookieHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rotateKey() {
	newcookie := securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	)

	atomic.StorePointer(&prevCookie, currentCookie)
	atomic.StorePointer(&currentCookie, unsafe.Pointer(newcookie))
}

func RotateKey(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			break
		case <-ticker.C:
		}

		rotateKey()
	}
}
