package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s) //// dot on a package — function inside a package (NOT OOP)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word] += 1
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}