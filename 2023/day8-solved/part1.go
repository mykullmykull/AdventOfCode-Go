package day8

import (
	"strings"
)

type node struct {
	left  string
	right string
}

func runA(input []string) int {
	directions, nodesMap := parseInput(input)
	return stepsToZZZ(directions, nodesMap)
}

func parseInput(input []string) (string, map[string]node) {
	directions := input[0]
	nodesMap := map[string]node{}
	for _, line := range input[1:] {
		name, node := parseLine(line)
		nodesMap[name] = node
	}
	return directions, nodesMap
}

func parseLine(line string) (string, node) {
	line = strings.Replace(line, "= (", "", 1)
	line = strings.Replace(line, ",", "", 1)
	line = strings.Replace(line, ")", "", 1)
	split := strings.Split(line, " ")

	return split[0], node{
		left:  split[1],
		right: split[2],
	}
}

func stepsToZZZ(directions string, nodesMap map[string]node) int {
	steps := 0
	i := 0
	location := "AAA"
	for {
		if directions[i] == 'L' {
			location = nodesMap[location].left
		} else {
			location = nodesMap[location].right
		}

		i = (i + 1) % len(directions)
		steps++

		if location == "ZZZ" {
			break
		}
	}
	return steps
}
