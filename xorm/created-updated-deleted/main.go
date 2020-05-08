package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Player struct {
	Id int64
	Name string
	Age int
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:12345@/test?charset=utf8")

	engine.Sync2(&Player{})
	engine.Insert(&Player{Name:"dj", Age:18})

	p1 := &Player{}
	engine.Where("name = ?", "dj").Get(p1)
	fmt.Println("after insert:", p1)
	time.Sleep(5 * time.Second)

	engine.Table(&Player{}).ID(p1.Id).Update(map[string]interface{}{"age":30})

	p2 := &Player{}
	engine.Where("name = ?", "dj").Get(p2)
	fmt.Println("after update:", p2)
	time.Sleep(5 * time.Second)

	engine.ID(p1.Id).Delete(&Player{})

	p3 := &Player{}
	engine.Where("name = ?", "dj").Unscoped().Get(p3)
	fmt.Println("after delete:", p3)
}