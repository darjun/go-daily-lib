package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/gocolly/colly/v2"
)

type Item struct {
	Id     string
	Width  int
	Height int
	Links  Links
}

type Links struct {
	Download string
}

func main() {
	c := colly.NewCollector(
		colly.Async(true),
	)
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `unsplash\.com`,
		RandomDelay:  500 * time.Millisecond,
		Parallelism:  12,
	})
	if err != nil {
		log.Fatal(err)
	}
	d := c.Clone()
	c.OnResponse(func(r *colly.Response) {
		var items []*Item
		json.Unmarshal(r.Body, &items)
		for _, item := range items {
			d.Visit(item.Links.Download)
		}
	})

	var count uint32
	d.OnResponse(func(r *colly.Response) {
		fileName := fmt.Sprintf("images/img%d.jpg", atomic.AddUint32(&count, 1))
		err := r.Save(fileName)
		if err != nil {
			fmt.Printf("saving %s failed:%v\n", fileName, err)
		} else {
			fmt.Printf("saving %s success\n", fileName)
		}
	})

	d.OnRequest(func(r *colly.Request) {
		fmt.Println("d visiting", r.URL)
	})
	d.OnError(func(r *colly.Response, err error) {
		fmt.Println("d error:", err, "url:", r.Request.URL)
	})

	for page := 1; page <= 10; page++ {
		c.Visit(fmt.Sprintf("https://unsplash.com/napi/photos?page=%d&per_page=12", page))
	}
	c.Wait()
	d.Wait()
}
