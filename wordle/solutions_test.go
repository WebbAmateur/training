package wordle

import "testing"

func TestSolutions(t *testing.T) {
	solutions, err := solutions()

	if err != nil {
		t.Errorf("Failure in solutions:%s", err)
	}

	if len(solutions) != SOLUTION_COUNT {
		t.Errorf("Solutions Expected:%d  Actual: %d", SOLUTION_COUNT, len(solutions))
	}
}
