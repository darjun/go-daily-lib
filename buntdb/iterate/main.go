package main

import (
	"fmt"
	"log"

	"github.com/tidwall/buntdb"
)

func main() {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *buntdb.Tx) error {
		data := map[string]string{
			"a": "apple",
			"b": "banana",
			"p": "pear",
			"o": "orange",
		}
		for key, value := range data {
			tx.Set(key, value, nil)
		}
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		var count int
		tx.Ascend("", func(key, value string) bool {
			fmt.Printf("key:%s value:%s\n", key, value)
			count++
			if count >= 3 {
				return false
			}
			return true
		})
		return nil
	})
}
