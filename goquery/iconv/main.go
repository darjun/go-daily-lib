package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

func SinaNewSurvival() {
	res, err := http.Get("http://news.sina.com.cn/society/netsurvival")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	utf8Body, err := iconv.NewReader(res.Body, "gb2312", "utf-8")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(utf8Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".title14 li").Each(func(i int, s *goquery.Selection) {
		content := s.Find("a").Text()
		time := s.Find("font").Text()
		fmt.Printf("%d: %s%s\n", i, content, time)
	})
}

func main() {
	SinaNewSurvival()
}
