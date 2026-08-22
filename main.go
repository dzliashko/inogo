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

func normalizeRefreshInterval(requested int) (int, bool) {

	if requested <= 0 {
		return 0, false
	}

	if requested < 5 {
		return 5, true
	} else if requested > 60 {
		return 60, true
	} else {
		return requested, true
	}

}

func buildRefreshPlan(
	enabled bool,
	requestedInterval int,
	maxAttempts int,
	attemptsUsed int,
) (int, string) {
	interval, intervalValid := normalizeRefreshInterval(requestedInterval)
	attempts, attemptsValid := calculateRemainingAttempts(maxAttempts, attemptsUsed)

	if !intervalValid || !attemptsValid {
		return 0, "invalid"
	}

	if !enabled {
		return interval, "disabled"
	}

	if attempts == 0 {
		return interval, "exhausted"
	}

	return interval, "ready"
}

func isRefreshCandidate(feedNumber int, pauseEvery int) bool {
	if feedNumber%pauseEvery == 0 {
		return false
	}
	return true
}

func countRefreshCandidates(totalFeeds int, pauseEvery int) (int, bool) {
	if totalFeeds < 0 || pauseEvery <= 0 {
		return 0, false
	}
	total := 0
	for i := 1; i <= totalFeeds; i++ {
		if isRefreshCandidate(i, pauseEvery) {
			total++
		}
	}
	return total, true
}

func lastFeedTitle(titles []string) (string, bool) {
	if len(titles) == 0 {
		return "", false
	}
	return titles[len(titles)-1], true
}

func sliceState(titles []string) string {
	if titles == nil {
		return "nil"
	}
	if len(titles) == 0 {
		return "empty"
	}
	return "non-empty"
}

func main() {
	sourceA := []string{"Go", "Postgres", "RSS"}
	alias := sourceA

	alias[0] = "Changed"

	fmt.Println("после изменения alias:")
	fmt.Println(sourceA)
	fmt.Println(alias)
	fmt.Println()

	fmt.Println("независимая копия:")
	sourceB := []string{"Go", "Postgres", "RSS"}
	destination := make([]string, len(sourceB))
	copied := copy(destination, sourceB)
	destination[0] = "Changed"
	fmt.Printf("copied %d\nsourceB: %v\ndestination: %v\n", copied, sourceB, destination)
	fmt.Println()

	fmt.Println("короткая копия:")
	destination2 := make([]string, 2)
	copied = copy(destination2, sourceB)
	fmt.Printf("copied %d\ndestination: %v\n", copied, destination2)
	fmt.Println()

	capacityOnly := make([]string, 0, 3)
	copied = copy(capacityOnly, sourceB)
	fmt.Printf("copied %d\nlen %d\ncap %d\ndestination: %v", copied, len(capacityOnly), cap(capacityOnly), capacityOnly)

}
