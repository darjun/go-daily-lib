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

	db.CreateIndex("names", "user:*:name", buntdb.IndexString)
	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:1:name", "tom", nil)
		tx.Set("user:2:name", "Randi", nil)
		tx.Set("user:3:name", "jane", nil)
		tx.Set("user:4:name", "Janet", nil)
		tx.Set("user:5:name", "Paula", nil)
		tx.Set("user:6:name", "peter", nil)
		tx.Set("user:7:name", "Terri", nil)
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("names", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return true
		})
		return nil
	})
}
