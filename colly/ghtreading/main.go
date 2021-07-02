package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Repository struct {
	Author  string
	Name    string
	Link    string
	Desc    string
	Lang    string
	Stars   int
	Forks   int
	Add     int
	BuiltBy []string
}

func main() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	repos := make([]*Repository, 0, 15)
	c.OnHTML(".Box .Box-row", func(e *colly.HTMLElement) {
		repo := &Repository{}

		// author & repository name
		authorRepoName := e.ChildText("h1.h3 > a")
		parts := strings.Split(authorRepoName, "/")
		repo.Author = strings.TrimSpace(parts[0])
		repo.Name = strings.TrimSpace(parts[1])

		// link
		repo.Link = e.Request.AbsoluteURL(e.ChildAttr("h1.h3 >a", "href"))

		// description
		repo.Desc = e.ChildText("p.pr-4")

		// language
		repo.Lang = strings.TrimSpace(e.ChildText("div.mt-2 > span.mr-3 > span[itemprop]"))

		// star & fork
		starForkStr := e.ChildText("div.mt-2 > a.mr-3")
		starForkStr = strings.Replace(strings.TrimSpace(starForkStr), ",", "", -1)
		parts = strings.Split(starForkStr, "\n")
		repo.Stars, _ = strconv.Atoi(strings.TrimSpace(parts[0]))
		repo.Forks, _ = strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))

		// add
		addStr := e.ChildText("div.mt-2 > span.float-sm-right")
		parts = strings.Split(addStr, " ")
		repo.Add, _ = strconv.Atoi(parts[0])

		// built by
		e.ForEach("div.mt-2 > span.mr-3  img[src]", func(index int, img *colly.HTMLElement) {
			repo.BuiltBy = append(repo.BuiltBy, img.Attr("src"))
		})

		repos = append(repos, repo)
	})

	c.Visit("https://github.com/trending")

	fmt.Printf("%d repositories\n", len(repos))
	fmt.Println("first repository:")
	for _, repo := range repos {
		fmt.Println("Author:", repo.Author)
		fmt.Println("Name:", repo.Name)
		break
	}
}
