package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"time"
)

var (
	logger *log.Logger
)

func WithLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("path:%s process start...\n", r.URL.Path)
		defer func() {
			logger.Printf("path:%s process end...\n", r.URL.Path)
		}()
		handler.ServeHTTP(w, r)
	})
}

func PanicRecover(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Println(string(debug.Stack()))
			}
		}()

		handler.ServeHTTP(w, r)
	})
}

func Metric(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			logger.Printf("path:%s elapsed:%fs\n", r.URL.Path, time.Since(start).Seconds())
		}()
		handler.ServeHTTP(w, r)
	})
}

func authenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			// no cookie
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		data, _ := base64.StdEncoding.DecodeString(cookie.Value)
		values, _ := url.ParseQuery(string(data))
		if values.Get("username") != "dj" && values.Get("password") != "handsome" {
			// failed
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func init() {
	logger = log.New(os.Stdout, "[goweb]", log.Lshortfile|log.LstdFlags)
}
