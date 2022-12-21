package filter

import (
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	simple := []string{START, END, START, END, START, END, START, END}
	filterSequence(t, simple, 0)

	offset := []string{END, START, END, START, END, START}
	filterSequence(t, offset, 0)
}

func filterSequence(t *testing.T, sequence []string, expectedError int) {
	actualError := Filter(sequence)
	if actualError != expectedError {
		t.Errorf("Sequence: %s  Expected:%d  Actual:%d", strings.Join(sequence, ","), expectedError, actualError)
	}
}
