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

func classifyUnreadCount(count int) string {
	if count < 0 {
		return "invalid"
	}
	if count == 0 {
		return "empty"
	}
	if count <= 9 {
		return "few"
	}
	return "many"
}

func canRefreshFeed(enabled bool, failures int) bool {
	if enabled && failures >= 0 && failures <= 2 {
		return true
	}
	return false
}

func shouldShowArticle(isRead bool, isStarred bool) bool {
	return !isRead || isStarred
}

func main() {
	fmt.Println(canRefreshFeed(true, 0))
	fmt.Println(canRefreshFeed(true, 2))
	fmt.Println(canRefreshFeed(true, 3))
	fmt.Println(canRefreshFeed(false, 0))
	fmt.Println(canRefreshFeed(true, -1))

	fmt.Println(shouldShowArticle(true, true))
	fmt.Println(shouldShowArticle(true, false))
	fmt.Println(shouldShowArticle(false, true))
	fmt.Println(shouldShowArticle(false, false))
}
