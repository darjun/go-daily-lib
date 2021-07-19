package main

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/gorilla/mux"
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
		time.Sleep(1 * time.Second)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	logger = log.New(os.Stdout, "[goweb]", log.Lshortfile|log.LstdFlags)

	r := mux.NewRouter()
	r.Use(PanicRecover, WithLogger, Metric)
	InitBooksRouter(r)
	InitMoviesRouter(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
