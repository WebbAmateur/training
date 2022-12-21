package wordle

import "regexp"

func filter(filter string, candidates []string) ([]string, error) {

	expression, err := regexp.Compile(filter)

	if err != nil {
		return nil, err
	}
	results := make([]string, 0, len(candidates))

	for _, candidate := range candidates {
		if expression.MatchString(candidate) {
			results = append(results, candidate)
		}
	}

	return results, nil
}
