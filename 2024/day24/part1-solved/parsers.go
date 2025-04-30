package day

import (
	"goAoc2024/utils"
	"strings"
)

const (
	AND = 0
	OR  = 1
	XOR = 2
)

type gate struct {
	w1 string
	w2 string
	op int
}

func parse(input []string) (map[string]int, map[string]gate, []string) {
	wireValues := map[string]int{}
	gates := map[string]gate{}
	unknowns := []string{}
	gatesLine := 0
	for x, line := range input {
		if line == "" {
			gatesLine = x + 1
			break
		}
		split := strings.Split(line, ": ")
		wire := split[0]
		value := utils.Atoi(split[1])
		wireValues[wire] = value
	}

	for _, line := range input[gatesLine:] {
		split := strings.Split(line, " -> ")
		out := split[1]
		split = strings.Split(split[0], " ")
		w1 := split[0]
		op := split[1]
		w2 := split[2]
		g := gate{w1, w2, getOp(op)}
		gates[out] = g
		for _, w := range []string{w1, w2, out} {
			if _, ok := wireValues[w]; !ok {
				wireValues[w] = -1
				unknowns = append(unknowns, w)
			}
		}
	}
	return wireValues, gates, unknowns
}

func getOp(op string) int {
	switch op {
	case "AND":
		return AND
	case "OR":
		return OR
	case "XOR":
		return XOR
	default:
		panic("Unknown op" + op)
	}
}
