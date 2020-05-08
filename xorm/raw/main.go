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

	querySql := "select * from user limit 1"
	reuslts, _ := engine.Query(querySql)
	for _, record := range reuslts {
		for key, val := range record {
			fmt.Println(key, string(val))
		}
	}

	updateSql := "update `user` set name=? where id=?"
	res, _ := engine.Exec(updateSql, "dj", 1)
	fmt.Println(res.RowsAffected())
}
