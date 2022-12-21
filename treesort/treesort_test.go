package treesort

import (
	"testing"
)

func TestSort(t *testing.T) {
	values := []int{50, 20, 30, 10, 44, 33, 22, 11, 67, 84, 92}

	root := TreeSort(values)

	sortedValues := TreeWalk(root, make([]int, 0))

	if len(values) != len(sortedValues) {
		t.Errorf("Length   Expected:%d Actual:%d", len(values), len(sortedValues))
	}

	
	current := sortedValues[0]
	
	for i := 1; i < len(sortedValues)-1; i++ {
		next := sortedValues[i]
		if current >= next {
			t.Errorf("%d is not < %d", current, next)
		}
		current = next
	}
}
