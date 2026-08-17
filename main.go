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
	titles := []string{"Go", "", "HTTP", ""}
	fmt.Println(countNonEmptyTitles(titles))
}
