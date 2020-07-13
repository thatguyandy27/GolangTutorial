package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	counts := map[string]int{}

	for _, word := range words {

		counts[word] += 1
	}

	fmt.Println(counts)
}
