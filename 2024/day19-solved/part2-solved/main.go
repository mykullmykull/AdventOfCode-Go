package day

import (
	"strings"
)

func main(input []string) int {
	towels := strings.Split(input[0], ", ")
	patterns := input[2:]
	validStacks := 0
	for _, p := range patterns {
		possibilities, _ := countPossibilities(p, "", towels, map[string]int{})
		println("validStacks:", validStacks, "pattern:", p, "adding", possibilities)
		validStacks += possibilities
	}
	return validStacks
}

func countPossibilities(goal string, pattern string, towels []string, cache map[string]int) (int, map[string]int) {
	if goal == "" {
		return 1, cache
	}
	if possibilities, ok := cache[pattern]; ok {
		return possibilities, cache
	}
	for _, t := range towels {
		possibilities := 0
		if len(t) <= len(goal) && t == goal[:len(t)] {
			newGoal := goal[len(t):]
			newPattern := pattern + t
			possibilities, cache = countPossibilities(newGoal, newPattern, towels, cache)
		}
		cache[pattern] += possibilities
	}
	return cache[pattern], cache
}
