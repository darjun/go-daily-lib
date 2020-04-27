package main

import (
	"fmt"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	key := []byte("name")
	val := []byte("dj")

	db.Update(func(tx *nutsdb.Tx) error {
		tx.Put("bucket1", key, val, 0)
		return nil
	})

	db.Update(func(tx *nutsdb.Tx) error {
		tx.Put("bucket2", key, val, 0)
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {
		e, _ := tx.Get("bucket1", key)
		fmt.Println("val1:", string(e.Value))

		e, _ = tx.Get("bucket2", key)
		fmt.Println("val2:", string(e.Value))
		return nil
	})
}
