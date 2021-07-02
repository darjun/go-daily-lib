package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	c1 := colly.NewCollector(
		colly.Async(true),
	)
	c1.SetRequestTimeout(time.Minute)
	c2 := c1.Clone()
	c3 := c1.Clone()

	c1.OnHTML("figure[itemProp] a[itemProp]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href == "" {
			return
		}

		c2.Visit(e.Request.AbsoluteURL(href))
	})

	c2.OnHTML("div._1g5Lu > img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if src == "" {
			return
		}

		c3.Visit(src)
	})

	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c1 visiting", r.URL)
	})

	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("c1 visiting", r.Request.URL, "failed:", err)
	})

	c2.OnRequest(func(r *colly.Request) {
		fmt.Println("c2 visiting", r.URL)
	})

	c2.OnError(func(r *colly.Response, err error) {
		fmt.Println("c2 visiting", r.Request.URL, "failed:", err)
	})

	var count uint32
	c3.OnResponse(func(r *colly.Response) {
		fileName := fmt.Sprintf("images/img%d.jpg", atomic.AddUint32(&count, 1))
		err := r.Save(fileName)
		if err != nil {
			fmt.Printf("saving %s failed:%v\n", fileName, err)
		} else {
			fmt.Printf("saving %s success\n", fileName)
		}
	})

	c3.OnRequest(func(r *colly.Request) {
		fmt.Println("c3 visiting", r.URL)
	})
	c3.OnError(func(r *colly.Response, err error) {
		fmt.Println("c3 error:", err)
	})

	c1.Visit("https://unsplash.com/")
	c1.Wait()
	c2.Wait()
	c3.Wait()
}
