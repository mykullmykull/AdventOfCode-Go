package day8

import (
	"fmt"
)

func runB(input []string) int {
	directions, nodeMaps := parseInput(input)
	startingLocations := getKeysEndingWithA(nodeMaps)
	zIntervals := findZIntervals(startingLocations, directions, nodeMaps)
	return findLeastCommonMultiple(zIntervals)
}

func getKeysEndingWithA(nodeMaps map[string]node) []string {
	keys := []string{}
	i := 0
	for k, _ := range nodeMaps {
		if k[len(k)-1] == 'A' {
			keys = append(keys, k)
		}
		i++
	}
	return keys
}

func findZIntervals(locations []string, directions string, nodesMap map[string]node) []int {
	zIntervals := make([]int, len(locations))
	for i, location := range locations {
		zIntervals[i] = findZInterval(location, directions, nodesMap)
	}
	return zIntervals
}

func findZInterval(location string, directions string, nodesMap map[string]node) int {
	intervals := []int{}
	steps := 0
	i := 0
	for {
		if directions[i] == 'L' {
			location = nodesMap[location].left
		} else {
			location = nodesMap[location].right
		}

		i = (i + 1) % len(directions)
		steps++

		if location[2] == 'Z' {
			intervals = append(intervals, steps)
		}

		if len(intervals) == 3 {
			if intervals[1] != 2*intervals[0] || intervals[2] != 3*intervals[0] {
				panic(fmt.Sprintf("starting at location %s we get different intervals: %v\n", location, intervals))
			}
			break
		}
	}
	return intervals[0]
}

func findLeastCommonMultiple(zIntervals []int) int {
	base := zIntervals[0]
	for _, interval := range zIntervals[1:] {
		gcd := getGreatestCommonDivisor(base, interval)
		base = base * interval / gcd
	}
	return base
}

func getGreatestCommonDivisor(a int, b int) int {
	max := a
	min := b

	for {
		if min == max {
			return max
		}

		if min > max {
			newMax := min
			min = max
			max = newMax
		}

		max = max - min
	}
}
