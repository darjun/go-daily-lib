package main

import (
	"fmt"

	"github.com/tidwall/buntdb"
)

func main() {
	db, _ := buntdb.Open(":memory:")
	defer db.Close()

	db.CreateIndex("first_name_age", "user:*", buntdb.IndexJSON("name.first"), buntdb.IndexJSON("age"))
	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:1", `{"name":{"first":"zhang","last":"san"},"age":18}`, nil)
		tx.Set("user:2", `{"name":{"first":"li","last":"si"},"age":27`, nil)
		tx.Set("user:3", `{"name":{"first":"wang","last":"wu"},"age":30}`, nil)
		tx.Set("user:4", `{"name":{"first":"sun","last":"qi"},"age":8}`, nil)
		tx.Set("user:5", `{"name":{"first":"li", "name":"dajun"},"age":20}`, nil)
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("first_name_age", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return true
		})
		return nil
	})
}
