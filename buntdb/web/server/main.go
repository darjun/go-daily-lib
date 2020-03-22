package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func init() {
	var err error
	db, err = buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
}

func response(w http.ResponseWriter, err error, data interface{}) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"error": err,
		"data":  data,
	})
	w.Write(bytes)
}

func set(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	value := r.FormValue("value")
	expire, _ := strconv.ParseBool(r.FormValue("expire"))
	ttl, _ := time.ParseDuration(r.FormValue("ttl"))

	var setOption *buntdb.SetOptions
	if expire && ttl > 0 {
		setOption = &buntdb.SetOptions{Expires: true, TTL: ttl}
	}

	err := db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, setOption)
		return err
	})

	response(w, err, nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")

	var value string
	err := db.View(func(tx *buntdb.Tx) error {
		var err error
		value, err = tx.Get(key)
		return err
	})

	response(w, err, value)
}

type Pair struct {
	Key   string
	Value string
}

func iterate(w http.ResponseWriter, r *http.Request) {
	index := r.FormValue("index")
	fmt.Println(index)

	var items []Pair
	err := db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend(index, func(key, value string) bool {
			fmt.Println(key, value)
			items = append(items, Pair{key, value})
			return true
		})
		return err
	})

	response(w, err, items)
}

func createIndex(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pattern := r.FormValue("pattern")
	less := buntdb.IndexString

	err := db.CreateIndex(name, pattern, less)
	response(w, err, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get", get)
	mux.HandleFunc("/set", set)
	mux.HandleFunc("/iterate", iterate)
	mux.HandleFunc("/create_index", createIndex)

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
