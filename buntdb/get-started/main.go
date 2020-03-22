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
		oldValue, replaced, err := tx.Set("testkey", "testvalue", nil)
		if err != nil {
			return err
		}

		fmt.Printf("old value:%q replaced:%t\n", oldValue, replaced)
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get("testkey")
		if err != nil {
			return err
		}

		fmt.Println("value is:", value)
		return nil
	})
}
