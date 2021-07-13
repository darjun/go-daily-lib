package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

type Middleware func(http.Handler) http.Handler

var (
	logger *log.Logger
)

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g)
}

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
		time.Sleep(1 * time.Second)
		handler.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

func main() {
	logger = log.New(os.Stdout, "goweb", log.Lshortfile|log.LstdFlags)

	middlewares := []Middleware{
		PanicRecover,
		WithLogger,
		Metric,
	}
	mux := http.NewServeMux()
	mux.Handle("/", applyMiddlewares(http.HandlerFunc(index), middlewares...))
	mux.Handle("/greeting", applyMiddlewares(greeting("welcome, dj"), middlewares...))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
