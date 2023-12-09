package day9

func runB(input []string) int {
	histories := parseInput(input)
	sum := 0
	for _, history := range histories {
		sum += getPreviousValue(history)
	}
	return sum
}

func getPreviousValue(history []int) int {
	differences := history
	firstDifferences := []int{history[0]}
	for {
		if isAllZeros(differences) {
			break
		}
		differences = extrapolateDifferences(differences)
		firstDifferences = append(firstDifferences, differences[0])
	}

	return fillInPreviousValues(firstDifferences)
}

func fillInPreviousValues(firsts []int) int {
	prevValue := 0
	for i := len(firsts) - 1; i > -1; i-- {
		first := firsts[i]
		prevValue = first - prevValue
	}
	return prevValue
}
