package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type draw struct {
	count int
	color string
}

func runB(input []string) int {
	sum := 0
	for _, line := range input {
		// fmt.Printf("line: %s\n", line)
		maxCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		draws := extractDraws(line)
		for _, draw := range draws {
			// fmt.Printf("  count: %d, color: %s, prevMax: %d\n", draw.count, draw.color, maxCubes[draw.color])
			if draw.count > maxCubes[draw.color] {
				maxCubes[draw.color] = draw.count
				// fmt.Printf("  new max: %d\n", maxCubes[draw.color])
			}

		}
		sum += maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
		// fmt.Printf("new sum: %d\n\n", sum)
	}
	return sum
}

func runA(input []string) int {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sumOfPossibleIds := 0
	for i, line := range input {
		gameId := i + 1
		if isPossible(line, maxCubes) {
			sumOfPossibleIds += gameId
		}
	}
	return sumOfPossibleIds
}

func isPossible(line string, maxCubes map[string]int) bool {
	draws := extractDraws(line)
	for _, draw := range draws {
		if draw.count > maxCubes[draw.color] {
			return false
		}
	}
	return true
}

func extractDraws(line string) []draw {
	withoutGameId := strings.Split(line, ": ")[1]
	withoutSemiColons := strings.Replace(withoutGameId, ";", ",", -1)
	cubeCounts := strings.Split(withoutSemiColons, ", ")
	countsAndColors := []draw{}
	for _, cubeCount := range cubeCounts {
		split := strings.Split(cubeCount, " ")
		color := split[1]
		count, err := strconv.Atoi(split[0])
		if err != nil {
			panic(fmt.Sprintf("couldn't convert %s to int from cubeCount %s\n", split[0], cubeCount))
		}

		countsAndColors = append(countsAndColors, draw{count, color})
	}

	return countsAndColors
}
