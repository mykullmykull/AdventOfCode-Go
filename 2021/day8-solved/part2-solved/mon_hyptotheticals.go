package day

import (
	"strings"
)

func (m *monitor) filterByHypothetical() {
	for x, wire := range allWires() {
		alreadySeen := map[string]bool{}
		for len(m.segs[x]) > 1 {
			// fmt.Printf("starting hypothetical with m.segs[x]: %s\n", m.segs[x])
			s := m.segs[x][0:1]
			if alreadySeen[s] {
				break
			}
			alreadySeen[s] = true
			m.segs[x] = m.segs[x][1:]
			if m.isPossible(string(wire), s) {
				m.segs[x] += s
			}
		}
	}
}

func (m *monitor) isPossible(wire, seg string) bool {
	for x, wStr := range m.wStrs {
		requiredSegs := requiredSegs(m.digits[x])
		// wStr requires seg, but lacks wire
		if strings.Contains(requiredSegs, seg) && !strings.Contains(wStr, wire) {
			return false
		}

		excludedSegs := excludedSegs(m.digits[x])
		// wStr has wire, but could never have seg
		if strings.Contains(excludedSegs, seg) && strings.Contains(wStr, wire) {
			return false
		}
	}
	return true
}

func allPossibleSegs(digits string) string {
	possibleSegs := map[string]bool{}
	for _, d := range digits {
		segs := digitToLitSegs(string(d))
		for _, seg := range segs {
			possibleSegs[string(seg)] = true
		}
	}
	allPossible := ""
	for seg := range possibleSegs {
		allPossible += seg
	}
	return allPossible
}

func requiredSegs(digits string) string {
	requiredSegs := allSegsStr()
	for _, d := range digits {
		dSegs := digitToLitSegs(string(d))
		dSegsStr := ""
		for _, seg := range dSegs {
			dSegsStr += string(seg)
		}
		for _, seg := range requiredSegs {
			if strings.Contains(dSegsStr, string(seg)) {
				continue
			}
			requiredSegs = strings.ReplaceAll(requiredSegs, string(seg), "")
		}
	}
	return requiredSegs
}

func excludedSegs(digits string) string {
	excludedSegs := allSegsStr()
	for _, d := range digits {
		dSegs := digitToLitSegs(string(d))
		for _, seg := range dSegs {
			excludedSegs = strings.ReplaceAll(excludedSegs, string(seg), "")
		}
	}
	return excludedSegs
}
