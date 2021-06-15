package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Repo struct {
	Name    string   // 仓库名
	Author  string   // 作者名
	Link    string   // 链接
	Desc    string   // 描述
	Lang    string   // 语言
	Stars   int      // 星数
	Forks   int      // fork 数
	Add     int      // 周期内新增
	BuiltBy []string // 贡献值 avatar img 链接
}

func getTrending(language, dateRange string) ([]*Repo, error) {
	repos := make([]*Repo, 0, 1)
	resp, err := http.Get(fmt.Sprintf("https://github.com/trending/%s?since=%s", language, dateRange))
	if err != nil {
		return repos, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return repos, err
	}

	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		titleSel := s.Find("h1 a")
		repo := &Repo{}
		authorSpanSel := titleSel.Find("span")
		repo.Author = strings.TrimSpace(strings.TrimRight(authorSpanSel.Text(), "/\n "))
		repo.Name = strings.TrimSpace(titleSel.Contents().Last().Text())
		relativeLink, _ := titleSel.Attr("href")
		repo.Link = "https://github.com/" + relativeLink
		repo.Desc = strings.TrimSpace(s.Find("p").Text())

		spanSel := s.Find("div>span")
		var langIndex int
		var addStarsIndex int
		var builtByIndex int
		if spanSel.Size() == 2 {
			langIndex = -1
			builtByIndex = 0
			addStarsIndex = 1
		} else {
			langIndex = 0
			builtByIndex = 1
			addStarsIndex = 2
		}
		if langIndex >= 0 {
			repo.Lang = strings.TrimSpace(spanSel.Eq(langIndex).Text())
		} else {
			repo.Lang = "unknown"
		}

		parts := strings.SplitN(strings.TrimSpace(spanSel.Eq(addStarsIndex).Text()), " ", 2)
		addStars, _ := strconv.Atoi(parts[0])
		repo.Add = addStars
		spanSel.Eq(builtByIndex).Find("a>img").Each(func(i int, img *goquery.Selection) {
			src, _ := img.Attr("src")
			repo.BuiltBy = append(repo.BuiltBy, src)
		})

		anchorSel := s.Find("div>a")
		starStr := strings.TrimSpace(anchorSel.Eq(-2).Text())
		star, _ := strconv.Atoi(strings.Replace(starStr, ",", "", -1))
		repo.Stars = star
		forkStr := strings.TrimSpace(anchorSel.Eq(-1).Text())
		fork, _ := strconv.Atoi(strings.Replace(forkStr, ",", "", -1))
		repo.Forks = fork

		repos = append(repos, repo)
	})
	return repos, nil
}
