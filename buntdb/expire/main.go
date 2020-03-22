package main

import (
	"fmt"
	"time"

	"github.com/tidwall/buntdb"
)

func main() {
	db, _ := buntdb.Open(":memory:")
	defer db.Close()

	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("testkey", "testvalue", &buntdb.SetOptions{Expires: true, TTL: time.Second})
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		value, _ := tx.Get("testkey")
		fmt.Println("value is:", value)
		return nil
	})

	time.Sleep(time.Second)

	db.View(func(tx *buntdb.Tx) error {
		value, _ := tx.Get("testkey")
		fmt.Println("value is:", value)
		return nil
	})
}
