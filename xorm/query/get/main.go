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
	has, _ := engine.ID(1).Get(user1)
	if has {
		fmt.Printf("user1:%v\n", user1)
	}

	user2 := &User{}
	has, _ = engine.Where("name=?", "dj").Get(user2)
	if has {
		fmt.Printf("user2:%v\n", user2)
	}

	user3 := &User{Id: 5}
	has, _ = engine.Get(user3)
	if has {
		fmt.Printf("user3:%v\n", user3)
	}

	user4 := &User{Name: "pipi"}
	has, _ = engine.Get(user4)
	if has {
		fmt.Printf("user4:%v\n", user4)
	}
}
