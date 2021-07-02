package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

func MyHeader(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("My-Header", "dj")
	})
}

func main() {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	MyHeader(c)

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("hello")
	})
	c.Visit("http://baidu.com")
}
