package day

import (
	"goAoc2024/utils"
	"strings"
)

func main(input []string) int {
	total := 0
	for _, line := range input {
		target, elements := parse(line)
		possibleResults := generatePossibleResults(target, elements)
		if reachedTarget(possibleResults, target) {
			total += target
		}
	}
	return total
}

func parse(line string) (int, []int) {
	split := strings.Split(line, ": ")
	target := utils.Atoi(split[0])
	elementStrs := strings.Split(split[1], " ")
	elements := []int{}
	for _, str := range elementStrs {
		elements = append(elements, utils.Atoi(str))
	}
	return target, elements
}

func generatePossibleResults(target int, elements []int) []int {
	possibleResults := []int{elements[0]}
	for _, e := range elements[1:] {
		newResults := []int{}
		for _, r := range possibleResults {
			sum := r + e
			product := r * e
			if sum <= target {
				newResults = append(newResults, sum)
			}
			if product <= target {
				newResults = append(newResults, product)
			}
		}
		possibleResults = newResults
	}
	return possibleResults
}

func reachedTarget(possibleResults []int, target int) bool {
	for _, r := range possibleResults {
		if r == target {
			return true
		}
	}
	return false
}
