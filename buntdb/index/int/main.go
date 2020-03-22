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

	db.CreateIndex("ages", "user:*:age", buntdb.IndexInt)
	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:1:age", "16", nil)
		tx.Set("user:2:age", "35", nil)
		tx.Set("user:3:age", "24", nil)
		tx.Set("user:4:age", "32", nil)
		tx.Set("user:5:age", "25", nil)
		tx.Set("user:6:age", "28", nil)
		tx.Set("user:7:age", "31", nil)
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("ages", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return true
		})
		return nil
	})
}
