package day9

import (
	"goAoc2023/utils"
	"strings"
)

func runA(input []string) int {
	histories := parseInput(input)
	sum := 0
	for _, history := range histories {
		sum += getNextValue(history)
	}
	return sum
}

func parseInput(input []string) [][]int {
	histories := [][]int{}
	for _, line := range input {
		historyStrs := strings.Split(line, " ")
		history := make([]int, len(historyStrs))
		for i, item := range historyStrs {
			history[i] = utils.Atoi(item)
		}
		histories = append(histories, history)
	}
	return histories
}

func getNextValue(history []int) int {
	differences := history
	lastDifferences := []int{history[len(history)-1]}
	for {
		if isAllZeros(differences) {
			break
		}
		differences = extrapolateDifferences(differences)
		lastDifferences = append(lastDifferences, differences[len(differences)-1])
	}

	nextValue := 0
	for _, lastDifference := range lastDifferences {
		nextValue += lastDifference
	}
	return nextValue
}

func isAllZeros(ints []int) bool {
	for _, n := range ints {
		if n != 0 {
			return false
		}
	}
	return true
}

func extrapolateDifferences(lastRow []int) []int {
	differences := make([]int, len(lastRow)-1)
	for i := 0; i < len(differences); i++ {
		differences[i] = lastRow[i+1] - lastRow[i]
	}
	return differences
}
