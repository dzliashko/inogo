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
	} else if count == 0 {
		return "empty"
	} else if count <= 9 {
		return "few"
	} else {
		return "many"
	}
}

func main() {
	fmt.Println(classifyUnreadCount(-1))
	fmt.Println(classifyUnreadCount(0))
	fmt.Println(classifyUnreadCount(1))
	fmt.Println(classifyUnreadCount(9))
	fmt.Println(classifyUnreadCount(10))

}
