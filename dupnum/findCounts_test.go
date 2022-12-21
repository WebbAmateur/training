package main

import "testing"

func TestFindCounts(t *testing.T) {

	nums := []int{1, 2, 3, 4, 3, 2, 1, 5, 4, 6}

	results := findCounts(nums)

	checkForError(t, results, 0, 0)
	checkForError(t, results, 1, 2)
	checkForError(t, results, 2, 2)
	checkForError(t, results, 3, 2)
	checkForError(t, results, 4, 2)
	checkForError(t, results, 5, 1)
	checkForError(t, results, 6, 1)
	checkForError(t, results, 7, 0)
	checkForError(t, results, 8, 0)
	checkForError(t, results, 9, 0)
}

func checkForError(t *testing.T, results []int, index int, expected int) {
	if int(results[index]) != expected {
		t.Errorf("At Index: %d   Expected:%d   Actual:%d", index, expected, results[index])
	}
}
