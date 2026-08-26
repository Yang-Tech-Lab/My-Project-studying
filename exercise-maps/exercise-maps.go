package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		counts[word]++
	}
	return counts
	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
