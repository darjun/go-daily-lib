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

	bucket := "list"
	key := []byte("userList")

	db.Update(func(tx *nutsdb.Tx) error {
		// 从头部依次插入多个值，注意顺序
		tx.LPush(bucket, key, []byte("user1"), []byte("user3"), []byte("user5"))
		// 当前list：user5, user3, user1

		// 从尾部依次插入多个值
		tx.RPush(bucket, key, []byte("user7"), []byte("user9"), []byte("user11"))
		// 当前list：user5, user3, user1, user7, user9, user11
		return nil
	})

	db.Update(func(tx *nutsdb.Tx) error {
		// 从头部删除一个值
		tx.LPop(bucket, key)
		// 当前list：user3, user1, user7, user9, user11

		// 从尾部删除一个值
		tx.RPop(bucket, key)
		// 当前list：user3, user1, user7, user9

		// 从头部删除两个值
		tx.LRem(bucket, key, 2)
		// 当前list：user7, user9
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {
		// 头部第一个值，user7
		b, _ := tx.LPeek(bucket, key)
		fmt.Println(string(b))

		// 长度
		l, _ := tx.LSize(bucket, key)
		fmt.Println(l)
		return nil
	})
}
