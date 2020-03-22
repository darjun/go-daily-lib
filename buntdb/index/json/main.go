package main

import (
	"fmt"

	"github.com/tidwall/buntdb"
)

func main() {
	db, _ := buntdb.Open(":memory:")
	defer db.Close()

	db.CreateIndex("first_name", "user:*", buntdb.IndexJSON("name.first"))
	db.CreateIndex("age", "user:*", buntdb.IndexJSON("age"))
	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:1", `{"name":{"first":"zhang","last":"san"},"age":18}`, nil)
		tx.Set("user:2", `{"name":{"first":"li","last":"si"},"age":27`, nil)
		tx.Set("user:3", `{"name":{"first":"wang","last":"wu"},"age":30}`, nil)
		tx.Set("user:4", `{"name":{"first":"sun","last":"qi"},"age":8}`, nil)
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		fmt.Println("Order by first name")
		tx.Ascend("first_name", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return true
		})

		fmt.Println("Order by age")
		tx.Ascend("age", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return true
		})

		fmt.Println("Order by age range 18-30")
		tx.AscendRange("age", `{"age":18}`, `{"age":30}`, func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return true
		})
		return nil
	})
}
