package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ISBN        string   `json:"isbn"`
	Name        string   `json:"name"`
	Authors     []string `json:"authors"`
	Press       string   `json:"press"`
	PublishedAt string   `json:"published_at"`
}

var (
	mapBooks map[string]*Book
	slcBooks []*Book
)

func init() {
	mapBooks = make(map[string]*Book)
	slcBooks = make([]*Book, 0, 1)

	data, err := ioutil.ReadFile("../data/books.json")
	if err != nil {
		log.Fatalf("failed to read books.json:%v", err)
	}

	err = json.Unmarshal(data, &slcBooks)
	if err != nil {
		log.Fatalf("failed to unmarshal books:%v", err)
	}

	for _, book := range slcBooks {
		mapBooks[book.ISBN] = book
	}
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(slcBooks)
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	book, ok := mapBooks[mux.Vars(r)["isbn"]]
	if !ok {
		http.NotFound(w, r)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(book)
}

func InitBooksRouter(r *mux.Router) {
	bs := r.PathPrefix("/books").Subrouter()
	bs.HandleFunc("/", BooksHandler)
	bs.HandleFunc("/{isbn}", BookHandler)
}
