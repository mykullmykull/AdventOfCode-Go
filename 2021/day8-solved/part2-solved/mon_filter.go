package day

import (
	"fmt"
	"strings"
)

func (m *monitor) eliminateDigits() {
	for x, wStr := range m.wStrs {
		digits := m.digits[x]
		maxY := len(digits)
		for y := 0; y < maxY; y++ {
			d := digits[0:1]
			digits = digits[1:]
			segs := digitToLitSegs(d)
			if len(segs) == len(wStr) {
				digits += d
				continue
			}
		}
		m.digits[x] = digits
	}
}

// eliminate segs wire does NOT belong to a wireStr with a possible digit with that segment
func (m *monitor) eliminateSegs() {
	allWires := allWires()
	for x := range m.segs {
		wire := allWires[x]
		for y, wStr := range m.wStrs {
			if !strings.Contains(wStr, string(wire)) {
				continue
			}
			digits := m.digits[y]
			if len(digits) != 1 {
				continue
			}
			d := digits[0]
			dSegs := digitToLitSegs(string(d))
			m.RemoveMissingSegs(x, dSegs)
		}
	}
}

func (m *monitor) RemoveMissingSegs(x int, dSegs string) {
	newSegs := m.segs[x]
	for _, seg := range m.segs[x] {
		if strings.Contains(dSegs, string(seg)) {
			continue
		}
		newSegs = strings.ReplaceAll(newSegs, string(seg), "")
	}
	m.segs[x] = newSegs
}

// a small group of wires have the same small number of possible segments
func (m *monitor) eliminateMatchedSegs() {
	for x, these := range m.segs {
		sameSegs := []int{x}
		for y, those := range m.segs {
			if y <= x {
				continue
			}
			if segsMatch(these, those) {
				sameSegs = append(sameSegs, y)
			}
		}
		if len(sameSegs) == len(these) {
			m.RemoveMatchedSegs(these, sameSegs)
		}
	}
}

func segsMatch(these, those string) bool {
	if len(these) != len(those) {
		return false
	}
	for _, seg := range these {
		if !strings.Contains(those, string(seg)) {
			return false
		}
	}
	return true
}

func (m *monitor) RemoveMatchedSegs(these string, sameSegs []int) {
	matchingIndexes := ""
	for _, index := range sameSegs {
		matchingIndexes += fmt.Sprintf("%d ", index)
	}

	for x := range m.segs {
		if strings.Contains(matchingIndexes, fmt.Sprintf("%d", x)) {
			continue
		}
		for _, seg := range these {
			m.segs[x] = strings.ReplaceAll(m.segs[x], string(seg), "")
		}
	}
}
