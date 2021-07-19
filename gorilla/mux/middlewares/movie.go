package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	IMDB        string `json:"imdb"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
	Duration    uint32 `json:"duration"`
	Lang        string `json:"lang"`
}

var (
	mapMovies map[string]*Movie
	slcMovies []*Movie
)

func init() {
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

func InitMoviesRouter(r *mux.Router) {
	ms := r.PathPrefix("/movies").Subrouter()
	ms.HandleFunc("/", MoviesHandler)
	ms.HandleFunc("/{imdb}", MovieHandler)
}
