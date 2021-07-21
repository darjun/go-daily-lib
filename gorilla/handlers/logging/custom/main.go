package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s", g)
}

func myLogFormatter(writer io.Writer, params handlers.LogFormatterParams) {
	var buf bytes.Buffer
	buf.WriteString(time.Now().Format("2006-01-02 15:04:05 -0700"))
	buf.WriteString(fmt.Sprintf(` "%s %s %s" `, params.Request.Method, params.URL.Path, params.Request.Proto))
	buf.WriteString(strconv.Itoa(params.StatusCode))
	buf.WriteByte('\n')

	writer.Write(buf.Bytes())
}

func Logging(handler http.Handler) http.Handler {
	return handlers.CustomLoggingHandler(os.Stdout, handler, myLogFormatter)
}

func main() {
	r := mux.NewRouter()
	r.Use(Logging)
	r.HandleFunc("/", index)
	r.Handle("/greeting/", greeting("dj"))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
