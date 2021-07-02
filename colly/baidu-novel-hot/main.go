package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Hot struct {
	Rank   string `selector:"a > div.index_1Ew5p"`
	Name   string `selector:"div.content_1YWBm > a.title_dIF3B"`
	Author string `selector:"div.content_1YWBm > div.intro_1l0wp:nth-child(2)"`
	Type   string `selector:"div.content_1YWBm > div.intro_1l0wp:nth-child(3)"`
	Desc   string `selector:"div.desc_3CTjT"`
}

func main() {
	c := colly.NewCollector()

	hots := make([]*Hot, 0, 10)
	c.OnHTML("div.category-wrap_iQLoo", func(e *colly.HTMLElement) {
		hot := &Hot{}

		err := e.Unmarshal(hot)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		hots = append(hots, hot)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response:", len(r.Body))
	})

	err := c.Visit("https://top.baidu.com/board?tab=novel")
	if err != nil {
		fmt.Println("Visit error:", err)
		return
	}

	fmt.Printf("%d hots\n", len(hots))
	for _, hot := range hots {
		fmt.Println("first hot:")
		fmt.Println("Rank:", hot.Rank)
		fmt.Println("Name:", hot.Name)
		fmt.Println("Author:", hot.Author)
		fmt.Println("Type:", hot.Type)
		fmt.Println("Desc:", hot.Desc)
		break
	}
}
