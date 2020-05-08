package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
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
	f, err := os.Create("sql.log")
	if err != nil {
		panic(err)
	}

	engine.SetLogger(log.NewSimpleLogger(f))
	engine.ShowSQL(true)

	user := &User{}
	engine.ID(1).Omit("created", "updated").Get(user)
	fmt.Printf("user:%v\n", user)
}
