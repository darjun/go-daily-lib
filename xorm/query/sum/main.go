package main

import (
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Sum struct {
	Id    int64
	Money int32
	Rate  float32
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:12345@/test?charset=utf8")
	engine.Sync2(&Sum{})

	var slice []*Sum
	for i := 0; i < 100; i++ {
		slice = append(slice, &Sum{
			Money: rand.Int31n(10000),
			Rate:  rand.Float32(),
		})
	}
	engine.Insert(&slice)

	totalMoney, _ := engine.SumInt(&Sum{}, "money")
	fmt.Println("total money:", totalMoney)

	totalRate, _ := engine.Sum(&Sum{}, "rate")
	fmt.Println("total rate:", totalRate)

	totals, _ := engine.Sums(&Sum{}, "money", "rate")
	fmt.Printf("total money:%f & total rate:%f", totals[0], totals[1])
}
