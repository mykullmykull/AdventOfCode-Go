package day

import "slices"

func main(input []string) int {
	// track slice of scores
	// loop through lines
	// if hasBadCloseer continue
	// get closers
	// calculate score
	// add score to slice
	// sort slice
	// return middle score
	scores := []int{}
	for _, line := range input {
		hasBadCloser, closers := findClosers(line)
		if hasBadCloser {
			continue
		}
		score := calculateScore(closers)
		scores = append(scores, score)
	}
	slices.Sort(scores)
	middleIndex := len(scores) / 2
	return scores[middleIndex]
}

func findClosers(line string) (bool, string) {
	expectedClosers := ""
	for _, char := range line {
		closer, found := closers[char]
		if found {
			expectedClosers = string(closer) + expectedClosers
			continue
		}

		if len(expectedClosers) == 0 || char != rune(expectedClosers[0]) {
			return true, ""
		}
		expectedClosers = expectedClosers[1:]
	}
	return false, expectedClosers
}

func calculateScore(closers string) int {
	score := 0
	for _, char := range closers {
		points, found := points[char]
		if !found {
			panic("Unexpected character: " + string(char))
		}
		score = (score * 5) + points
	}
	println("closers: ", closers, " score: ", score)
	return score
}
