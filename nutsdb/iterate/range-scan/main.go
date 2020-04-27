package main

import (
	"fmt"
	"strconv"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	bucket := "user_list"
	prefix := "user_"
	db.Update(func(tx *nutsdb.Tx) error {
		for i := 1; i <= 300; i++ {
			key := []byte(prefix + strconv.FormatInt(int64(i), 10))
			val := []byte("dj" + strconv.FormatInt(int64(i), 10))
			tx.Put(bucket, key, val, 0)
		}
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {
		lbound := []byte("user_100")
		ubound := []byte("user_199")
		entries, _ := tx.RangeScan(bucket, lbound, ubound)
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		return nil
	})
}
