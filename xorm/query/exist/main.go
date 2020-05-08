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
	has, _ := engine.ID(1).Exist(user1)
	if has {
		fmt.Println("user with id=1 exist")
	} else {
		fmt.Println("user with id=1 not exist")
	}

	user2 := &User{}
	has, _ = engine.Where("name=?", "dj2").Get(user2)
	if has {
		fmt.Println("user with name=dj2 exist")
	} else {
		fmt.Println("user with name=dj2 not exist")
	}
}
