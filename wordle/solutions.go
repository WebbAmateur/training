package wordle

import (
	"bufio"
	"os"
)

const SOLUTION_COUNT int = 2309

func solutions() ([]string, error) {

	solutions := make([]string, 0, 2309)

	readFile, err := os.Open("solutions.txt")

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		solutions = append(solutions, fileScanner.Text())
	}

	return solutions, nil
}
