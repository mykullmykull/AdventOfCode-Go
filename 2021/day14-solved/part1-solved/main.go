package day

import "math"

func main(input []string) int {
	pairCount, evolutionRules := parse(input)
	for step := 0; step < 10; step++ {
		pairCount = evolve(pairCount, evolutionRules)
	}
	first := input[0][0]
	last := input[0][len(input[0])-1]

	letterCount := countLetters(pairCount, first, last)
	return mostMinusLeast(letterCount)
}

func evolve(pairCount map[string]int, evolutionRules map[string][2]string) map[string]int {
	newPairCount := make(map[string]int)
	for pair, count := range pairCount {
		rule := evolutionRules[pair]
		newPairCount[rule[0]] += count
		newPairCount[rule[1]] += count
	}
	return newPairCount
}

func countLetters(pairCount map[string]int, first byte, last byte) map[byte]int {
	letterCount := make(map[byte]int)
	for pair, count := range pairCount {
		letterCount[pair[0]] += count
		letterCount[pair[1]] += count
	}
	letterCount[first]++ // Count the first letter
	letterCount[last]++  // Count the last letter

	for b, count := range letterCount {
		letterCount[b] = count / 2
	}
	return letterCount
}

func mostMinusLeast(letterCount map[byte]int) int {
	max := 0
	min := math.MaxInt
	for _, count := range letterCount {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return max - min
}
