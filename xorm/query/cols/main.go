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

	user1 := &User{}
	engine.ID(1).Cols("id", "name", "age").Get(user1)
	fmt.Printf("user1:%v\n", user1)

	user2 := &User{Name: "pipi"}
	engine.Omit("created", "updated").Get(user2)
	fmt.Printf("user2:%v\n", user2)
}
