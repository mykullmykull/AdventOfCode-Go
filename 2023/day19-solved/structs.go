package day19

import "fmt"

type span struct {
	min int
	max int
}

type partSpan struct {
	x            span
	m            span
	a            span
	s            span
	atPileName   string
	nextPileName string
}

func (s span) totalPossibilities() int {
	diff := s.max - s.min + 1
	if diff <= 0 {
		return 0
	}
	return diff
}

func (s span) restrictBy(r rule, negate bool) span {
	if r.attribute == "" {
		return s
	}

	if r.operator == ">" {
		if negate {
			s.max = min(s.max, r.value)
		} else {
			s.min = max(s.min, r.value+1)
		}
	} else {
		if negate {
			s.min = max(s.min, r.value)
		} else {
			s.max = min(s.max, r.value-1)
		}
	}

	return s
}

func (ps partSpan) totalPossibilities() int {
	return ps.x.totalPossibilities() * ps.m.totalPossibilities() * ps.a.totalPossibilities() * ps.s.totalPossibilities()
}

func (ps partSpan) restrictBy(r rule, negate bool) partSpan {
	if r.attribute == "" {
		return ps
	}

	switch r.attribute {
	case "x":
		ps.x = ps.x.restrictBy(r, negate)
	case "m":
		ps.m = ps.m.restrictBy(r, negate)
	case "a":
		ps.a = ps.a.restrictBy(r, negate)
	case "s":
		ps.s = ps.s.restrictBy(r, negate)
	default:
		panic(fmt.Sprintf("unknown attribute: %s", r.attribute))
	}

	return ps
}

func (ps partSpan) toString() string {
	return fmt.Sprintf("x: %v, m: %v, a: %v, s: %v, at: %s, next: %s", ps.x, ps.m, ps.a, ps.s, ps.atPileName, ps.nextPileName)
}

func (p pile) targetedPartSpans(s partSpan, pileName string) []partSpan {
	debug := false
	partSpans := []partSpan{}
	for i, r := range p.rules {
		// log(fmt.Sprintf("  pile: %s, rule #%d: %v", pileName, i, r), debug)
		if r.destination == s.nextPileName {
			newSpan := s.restrictBy(r, false)
			log(fmt.Sprintf("    restricted by rule: %v, new partSpan: %v", r, newSpan), debug)
			for j := 0; j < i; j++ {
				newSpan = newSpan.restrictBy(p.rules[j], true)
				log(fmt.Sprintf("    restricted by negating rule: %v, new partSpan: %v", p.rules[j], newSpan), debug)
			}
			newSpan.atPileName = s.nextPileName
			newSpan.nextPileName = pileName
			log(fmt.Sprintf("  rule: %v, new partSpan: %v", r, newSpan), debug)
			partSpans = append(partSpans, newSpan)
		}
	}
	return partSpans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------------------- Part 1 -------------------------------------//

type pile struct { // in practice this will be a map[string]Pile to optimize searching by name
	rules []rule
}

type rule struct {
	attribute   string // final and default rule of each pile will have attribute ""
	operator    string
	value       int
	destination string
}

type part struct {
	x        int
	m        int
	a        int
	s        int
	pileName string
}

func (p part) getFinalDestination(piles map[string]pile) part {
	debug := false
	for {
		log(fmt.Sprintf("  updated part: %v", p), debug)
		if p.pileName == "A" || p.pileName == "R" {
			return p
		}
		p.pileName = p.getNextPileName(piles)
	}
}

func (p part) getNextPileName(piles map[string]pile) string {
	location := piles[p.pileName]
	for _, r := range location.rules {
		if p.matches(r) {
			return r.destination
		}
	}
	panic(fmt.Sprintf("no rule matches part %v", p))
}

func (p part) matches(r rule) bool {
	if r.attribute == "" {
		return true
	}
	switch r.attribute + r.operator {
	case "x>":
		return p.x > r.value
	case "x<":
		return p.x < r.value
	case "m>":
		return p.m > r.value
	case "m<":
		return p.m < r.value
	case "a>":
		return p.a > r.value
	case "a<":
		return p.a < r.value
	case "s>":
		return p.s > r.value
	case "s<":
		return p.s < r.value
	default:
		panic(fmt.Sprintf("unknown attribute/operator: %s%s", r.attribute, r.operator))
	}
}

func (p part) totalAttributes() int {
	return p.x + p.m + p.a + p.s
}
