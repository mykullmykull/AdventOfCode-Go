package day

func main(input []string) int {
	score := 0
	for _, line := range input {
		firstIllegalChar := findFirstIllegalChar(line)
		if firstIllegalChar == '0' {
			continue
		}

		points, found := points[firstIllegalChar]
		if !found {
			panic("Unexpected character: " + string(firstIllegalChar))
		}
		score += points
	}
	return score
}

func findFirstIllegalChar(line string) rune {
	expectedClosers := []rune{}
	for _, char := range line {
		closer, found := closers[char]
		if found {
			expectedClosers = append([]rune{closer}, expectedClosers...)
			continue
		}

		if len(expectedClosers) == 0 || char != expectedClosers[0] {
			return char
		}
		expectedClosers = expectedClosers[1:]
	}
	return '0'
}
