package day

import (
	"strings"
)

func main(input []string) int {
	towels := strings.Split(input[0], ", ")
	patterns := input[2:]
	possiblePatternsCount := 0
	for _, p := range patterns {
		if isPossible(p, towels) {
			possiblePatternsCount++
		}
	}
	return possiblePatternsCount
}

func isPossible(goal string, towels []string) bool {
	println()
	println("is pattern possible?", goal)
	seen := map[string]bool{}
	patterns := []string{""}
	seen[""] = true
	for len(patterns) > 0 {
		p := patterns[0]
		println("  patterns:", len(patterns), "pattern:", p)
		patterns = patterns[1:]
		for _, t := range towels {
			newPattern := p + t
			if newPattern == goal {
				println("    perfect")
				return true
			}
			if len(newPattern) >= len(goal) || seen[newPattern] {
				continue
			}
			seen[newPattern] = true
			if newPattern == goal[:len(newPattern)] {
				patterns = append(patterns, newPattern)
			}
		}
	}
	return false
}

func makePattern(indexes []int, towels []string) string {
	pattern := ""
	for _, i := range indexes {
		pattern += towels[i]
	}
	return pattern
}

func increment(indexes []int, max int) []int {
	if len(indexes) == 0 {
		return indexes
	}
	lastX := len(indexes) - 1
	indexes[lastX]++
	if indexes[lastX] >= max {
		indexes = indexes[:len(indexes)-1]
		return increment(indexes, max)
	}
	return indexes
}
