package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Initialize a place to store the numbers
	nums := make([]int, 0)

	// Create a slice of valid integers [0..9]
	for i := 1; i < len(os.Args); i++ {

		value64, err := strconv.ParseInt(os.Args[i], 10, 8)

		if err == nil && value64 >= 0 && value64 <= 9 {
			nums = append(nums, int(value64))
		}
	}

	counts := findCounts(nums)

	for index := range counts {
		if counts[index] > 1 {
			fmt.Printf("Count for %d: = %d\n", index, counts[index])
		}
	}
}

// findCounts return the number of time a number [0..9] appears in num
func findCounts(nums []int) []int {
	// Initialize a slice for the counts
	// Automatically initialized to 0
	result := make([]int, 10)

	for _, num := range nums {
		result[num] = result[num] + 1
	}

	return result
}
