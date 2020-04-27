package main

import (
	"fmt"
	"log"
	"time"

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
		tx.Put("", key, val, 10)
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {
		e, _ := tx.Get("", key)
		fmt.Println(string(e.Value))
		return nil
	})

	time.Sleep(10 * time.Second)

	db.View(func(tx *nutsdb.Tx) error {
		e, err := tx.Get("", key)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(e.Value))
		}
		return nil
	})
}
