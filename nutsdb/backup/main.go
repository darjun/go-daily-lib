package main

import (
	"fmt"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)

	key := []byte("name")
	val := []byte("dj")
	db.Update(func(tx *nutsdb.Tx) error {
		tx.Put("", key, val, 0)
		return nil
	})

	db.Backup("./backup")
	db.Close()

	opt.Dir = "./backup"
	backupDB, _ := nutsdb.Open(opt)
	backupDB.View(func(tx *nutsdb.Tx) error {
		e, _ := tx.Get("", key)
		fmt.Println(string(e.Value))
		return nil
	})
}
