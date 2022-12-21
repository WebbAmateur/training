package filter

var START string = "Start"
var END string = "End"

func Filter(log []string) int {
	errorCount := 0
	var index int = 0

	// Find First Start
	index = findNextStart(log, index)

	for index <= len(log)-2 {

		// log[index] == START

		// Expect END at index +1
		expectedEnd := log[index+1]

		if expectedEnd != END {
			errorCount++
			index = findNextStart(log, index+1)
		} else {
			// expectedEnd == END
			nextIndex := findNextStart(log, index+2)
			if nextIndex != index+2 {
				errorCount++
			}
			index = nextIndex
		}
	}

	return errorCount
}

// findNextStart begins at index in log and returns the index of the next START in the range [index .. len(log) -2]
// findNextStart returns len(log)  if no start is found
func findNextStart(log []string, index int) int {
	for loc := index; loc <= len(log)-1; loc++ {
		if log[loc] == START {
			return loc
		}
	}

	return len(log)
}
