package day

import (
	"goAoc2021/utils"
	"math"
	"strings"
)

type monitor struct {
	unsolved int
	wStrs    []string
	digits   []string // e.g. 0123, index matches wStrs index
	segs     []string // e.g. ABCD, index matches wStrs index
}

func (m *monitor) init(line string) {
	m.unsolved = math.MaxInt
	line = strings.ReplaceAll(line, " | ", " ")
	m.wStrs = strings.Split(line, " ")

	m.digits = make([]string, len(m.wStrs))
	for index := range m.wStrs {
		m.digits[index] = allDigitsStr()
	}

	m.segs = make([]string, len(allWires()))
	for index := range m.segs {
		m.segs[index] = allSegsStr()
	}

}

// isSimpler checks if the number of unsolved possibilities is less than previously
func (m *monitor) isSimpler() bool {
	newUnsolved := 0
	for _, d := range m.digits {
		newUnsolved += len(d)
	}
	for _, s := range m.segs {
		newUnsolved += len(s)
	}
	isSimpler := newUnsolved < m.unsolved
	m.unsolved = newUnsolved
	return isSimpler
}

// check if all wire strings and wires are solved
func (m *monitor) isSolvable(line string) bool {
	for _, segs := range m.segs {
		if len(segs) > 1 {
			return false
		}
	}
	return true
}

// solve decodes the output value from the solved wires
func (m monitor) solve(line string) int {
	output := strings.Split(line, " | ")[1]
	for x, wire := range allWires() {
		seg := m.segs[x][0]
		output = strings.ReplaceAll(output, string(wire), string(seg))
	}
	solvedLength := len(strings.Split(output, " "))*2 - 1
	for len(output) > solvedLength {
		segs := getLongestSegs(output)
		d := getDigit(segs)
		output = strings.ReplaceAll(output, segs, d)
	}
	output = strings.ReplaceAll(output, " ", "")
	return utils.Atoi(output)
}

func getDigit(segs string) string {
	for x, dSegs := range allLitSegs() {
		if segsMatch(segs, dSegs) {
			return allDigits()[x]
		}
	}
	return ""
}

func getLongestSegs(output string) string {
	longest := 0
	longestSegs := ""
	outputLitSegs := strings.Split(output, " ")
	for _, segs := range outputLitSegs {
		if len(segs) > longest {
			longest = len(segs)
			longestSegs = segs
		}
	}
	return longestSegs
}
