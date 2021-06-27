package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Repository struct {
	ID              int        `json:"id"`
	NodeID          string     `json:"node_id"`
	Name            string     `json:"name"`
	FullName        string     `json:"full_name"`
	Owner           *Developer `json:"owner"`
	Private         bool       `json:"private"`
	Description     string     `json:"description"`
	Fork            bool       `json:"fork"`
	Language        string     `json:"language"`
	ForksCount      int        `json:"forks_count"`
	StargazersCount int        `json:"stargazers_count"`
	WatchersCount   int        `json:"watchers_count"`
	OpenIssuesCount int        `json:"open_issues_count"`
}

type Developer struct {
	Login      string `json:"login"`
	ID         int    `json:"id"`
	NodeID     string `json:"node_id"`
	AvatarURL  string `json:"avatar_url"`
	GravatarID string `json:"gravatar_id"`
	Type       string `json:"type"`
	SiteAdmin  bool   `json:"site_admin"`
}

func main() {
	client := resty.New()

	var result []*Repository
	client.R().
		SetAuthToken("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").
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetQueryParams(map[string]string{
			"per_page":  "3",
			"page":      "1",
			"sort":      "created",
			"direction": "asc",
		}).
		SetPathParams(map[string]string{
			"org": "golang",
		}).
		SetResult(&result).
		Get("https://api.github.com/orgs/{org}/repos")

	for i, repo := range result {
		fmt.Printf("repo%d: name:%s stars:%d forks:%d\n", i+1, repo.Name, repo.StargazersCount, repo.ForksCount)
	}
}
