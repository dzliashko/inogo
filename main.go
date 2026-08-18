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

func countRefreshAttempts(maxAttempts int, succeedsOn int) int {
	attempt := 1
	if maxAttempts <= 0 {
		return 0
	}

	for attempt < maxAttempts {
		if attempt == succeedsOn {
			break
		}
		attempt++
	}
	return attempt
}

func remainingRefreshAttempts(maxAttempts int, attemptsUsed int) int {
	if maxAttempts <= 0 || attemptsUsed < 0 || maxAttempts <= attemptsUsed {
		return 0
	}
	return maxAttempts - attemptsUsed
}

func canAttemptRefresh(enabled bool, maxAttempts int, attemptsUsed int) bool {
	remaining := remainingRefreshAttempts(maxAttempts, attemptsUsed)
	if enabled && remaining > 0 {
		return true
	}
	return false
}

func calculateRemainingAttempts(maxAttempts int, attemptsUsed int) (int, bool) {
	if maxAttempts <= 0 || attemptsUsed < 0 || attemptsUsed > maxAttempts {
		return 0, false
	}
	if attemptsUsed == maxAttempts {
		return 0, true
	}
	return maxAttempts - attemptsUsed, true
}

func refreshAvailability(maxAttempts int, attemptsUsed int) string {
	remaining, isDataCorrect := calculateRemainingAttempts(maxAttempts, attemptsUsed)
	if remaining == 0 && !isDataCorrect {
		return "invalid"
	}
	if remaining == 0 && isDataCorrect {
		return "exhausted"
	}
	return "available"
}

func main() {
	fmt.Println(refreshAvailability(3, 0))
	fmt.Println(refreshAvailability(3, 2))
	fmt.Println(refreshAvailability(3, 3))
	fmt.Println(refreshAvailability(3, 5))
	fmt.Println(refreshAvailability(0, 0))
	fmt.Println(refreshAvailability(3, -1))
}
