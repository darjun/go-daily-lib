package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:12345@/test?charset=utf8")

	rows, _ := engine.Where("age > ? and age < ?", 12, 30).Rows(&User{})
	defer rows.Close()

	u := &User{}
	for rows.Next() {
		rows.Scan(u)

		fmt.Println(u)
	}
}
