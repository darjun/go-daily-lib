package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
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

	utf8Body, err := DecodeHTMLBody(res.Body, "")
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

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, _ := charset.DetermineEncoding(data, ""); len(name) != 0 {
			return name
		}
	}

	return "utf-8"
}

func DecodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
	if charset == "" {
		charset = detectContentCharset(body)
	}

	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}

	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}

	return body, nil
}
