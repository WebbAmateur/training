package wordle

import (
	"fmt"
	"testing"
)

func TestRegex(t *testing.T) {
	//arose := "[^aos][^aros][^aos][^aos][^aose]"
	//finer := "[in][fn][fi]er"
	endInO := "a...o"

	solutions, err := solutions()

	if err != nil {
		t.Errorf("Failure in solutions:%s", err)
	}

	testExpression(t, solutions, endInO)
	//testExpression(t, solutions, arose)

	//testExpression(t, solutions, finer)

	//known := "w[^arsekn]o[^arsekn][^arsekn]"

	//testExpression(t, solutions, known)

}

func testExpression(t *testing.T, solutions []string, expression string) {
	results, err := filter(expression, solutions)

	if err != nil {
		t.Errorf("Failure in filter():%s", err)
	}

	fmt.Printf("%d Candidates for %s\n", len(results), expression)
	for _, result := range results {
		fmt.Println(result)
	}

}
