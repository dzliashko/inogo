package main

import "fmt"

func countNonEmptyTitles(titles []string) int {
	total := 0
	for _, title := range titles {
		if title != "" {
			total++
		}
	}
	return total
}

func main() {
	var feed string
	var articles int
	var isUpdated bool

	fmt.Printf("Feed = %s, Articles = %d, Updated = %t\n", feed, articles, isUpdated)
	feed = "Go Blog"
	articles = 12
	isUpdated = true

	url := "https://ru.hexlet.io/blog.rss"
	const maxArticles = 1000
	fmt.Printf("Feed = %s, Articles = %d, Updated = %t, URL = %s, MaxArticles = %d\n", feed, articles, isUpdated, url, maxArticles)

}
