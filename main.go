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

func feedStatus(status string) string {
	switch status {
	case "active":
		return "refresh enabled"
	case "paused":
		return "refresh paused"
	case "failed":
		return "refresh failed"
	default:
		return "unknown state"
	}
}

func countProcessableArticles(total int) int {
	count := 0
	if total <= 0 {
		return 0
	}
	for i := 1; i <= total; i++ {
		if i > 10 {
			break
		}
		if i%3 == 0 {
			continue
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(countProcessableArticles(-1))
	fmt.Println(countProcessableArticles(0))
	fmt.Println(countProcessableArticles(1))
	fmt.Println(countProcessableArticles(3))
	fmt.Println(countProcessableArticles(10))
	fmt.Println(countProcessableArticles(20))
}
