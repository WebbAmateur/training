package main

import "testing"

func TestMapFindCounts(t *testing.T) {

	chars := []string{"a", "b", "c", "d", "a", "b", "d", "e", "f"}

	results := mapFindCounts(chars)

	checkForError(t, results, "a", 2)
	checkForError(t, results, "b", 2)
	checkForError(t, results, "c", 1)
	checkForError(t, results, "d", 1)
	checkForError(t, results, "e", 1)
	checkForError(t, results, "f", 1)
	checkForError(t, results, "g", 0)

}

func checkForError(t *testing.T, results map[string]int, key string, expected int) {
	value := results[key]

	if value != expected {
		t.Errorf("At Key: %s   Expected:%d   Actual:%d", key, expected, results[key])
	}
}
