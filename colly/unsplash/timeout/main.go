package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
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
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `unsplash\.com`,
		Delay:        1 * time.Second,
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
		log.Printf("d visiting:%s\n", r.URL)
	})
	d.OnError(func(r *colly.Response, err error) {
		log.Printf("d error:%v url:%s\n", err, r.Request.URL)
	})

	for page := 1; page <= 10; page++ {
		c.Visit(fmt.Sprintf("https://unsplash.com/napi/photos?page=%d&per_page=12", page))
	}
	c.Wait()
	d.Wait()
}
