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

func main() {
	fmt.Println(classifyUnreadCount(-1))
	fmt.Println(classifyUnreadCount(0))
	fmt.Println(classifyUnreadCount(1))
	fmt.Println(classifyUnreadCount(9))
	fmt.Println(classifyUnreadCount(10))

}
