package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialize a place to store the nchars
	chars := make([]string, 0)

	// Create a slice of valid chars [a..z]
	for i := 1; i < len(os.Args); i++ {

		current := os.Args[i]

		if len(current) == 1 {
			chars = append(chars, current)
		}
	}

	counts := mapFindCounts(chars)

	for key, value := range counts {
		if value > 1 {
			fmt.Printf("Count for %s: = %d\n", key, value)
		}
	}
}

// findCounts return the number of time a char [a..z] appears in chars
func mapFindCounts(chars []string) map[string]int {
	// Initialize a slice for the counts
	// Automatically initialized to 0
	result := make(map[string]int)

	return result
}
