package main

import (
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
	engine.ID(1).Update(&User{Name: "ldj"})
	engine.ID(1).Cols("name", "age").Update(&User{Name: "dj"})

	engine.Table(&User{}).ID(1).Update(map[string]interface{}{"age": 18})
}
