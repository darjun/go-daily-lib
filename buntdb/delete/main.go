package main

import (
	"fmt"
	"strconv"

	"github.com/tidwall/buntdb"
)

func main() {
	db, _ := buntdb.Open(":memory:")
	defer db.Close()

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

	db.Update(func(tx *buntdb.Tx) error {
		deleteKeys := make([]string, 0)
		tx.Ascend("", func(key, value string) bool {
			age, _ := strconv.ParseUint(value, 10, 64)
			if age >= 30 {
				deleteKeys = append(deleteKeys, key)
			}
			return true
		})

		for _, key := range deleteKeys {
			tx.Delete(key)
		}
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return true
		})
		return nil
	})
}
