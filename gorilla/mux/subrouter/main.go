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

type Movie struct {
	IMDB        string `json:"imdb"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
	Duration    uint32 `json:"duration"`
	Lang        string `json:"lang"`
}

var (
	mapBooks map[string]*Book
	slcBooks []*Book

	mapMovies map[string]*Movie
	slcMovies []*Movie
)

func initBooks() {
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

func initMovies() {
	mapMovies = make(map[string]*Movie)
	slcMovies = make([]*Movie, 0, 1)

	data, err := ioutil.ReadFile("../data/movies.json")
	if err != nil {
		log.Fatalf("failed to read movies.json:%v", err)
	}

	err = json.Unmarshal(data, &slcMovies)
	if err != nil {
		log.Fatalf("failed to unmarshal movies:%v", err)
	}

	for _, movie := range slcMovies {
		mapMovies[movie.IMDB] = movie
	}
}

func init() {
	initBooks()
	initMovies()
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

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(slcMovies)
}

func MovieHandler(w http.ResponseWriter, r *http.Request) {
	movie, ok := mapMovies[mux.Vars(r)["imdb"]]
	if !ok {
		http.NotFound(w, r)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(movie)
}

func main() {
	r := mux.NewRouter()
	bs := r.PathPrefix("/books").Subrouter()
	bs.HandleFunc("/", BooksHandler)
	bs.HandleFunc("/{isbn}", BookHandler)

	ms := r.PathPrefix("/movies").Subrouter()
	ms.HandleFunc("/", MoviesHandler)
	ms.HandleFunc("/{imdb}", MovieHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
