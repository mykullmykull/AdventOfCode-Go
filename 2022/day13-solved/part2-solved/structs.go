package day

type pair struct {
	ltStr       string
	rtStr       string
	lt          set
	rt          set
	isLtSmaller int
}

type set struct {
	n       int
	subsets []subset
}

type subset struct {
	s set
}

const (
	yes   = 1
	maybe = 0
	no    = -1
)

func (s set) insertInOrder(sets []set) []set {
	newSets := []set{}
	for x, s2 := range sets {
		p := pair{lt: s, rt: s2}
		p = p.checkIfLtIsSmaller("  ")
		if p.isLtSmaller == yes {
			newSets = append(newSets, sets[:x]...)
			newSets = append(newSets, s)
			newSets = append(newSets, sets[x:]...)
			return newSets
		}
	}
	return append(sets, s)
}

func (s set) getIndex(sets []set) int {
	for x, s2 := range sets {
		if len(s2.subsets) == 0 {
			continue
		}

		if s.subsets[0].s.n == s2.subsets[0].s.n {
			return x + 1
		}
	}
	return -1
}

func (s set) defaultSet() set {
	s.n = -1
	s.subsets = []subset{}
	return s
}

func (s set) isEmpty() bool {
	return !s.isInt() && len(s.subsets) == 0
}

func (s set) isInt() bool {
	return s.n > -1
}

func (p pair) isInOrder() bool {
	p = p.checkIfLtIsSmaller("  ")
	if p.isLtSmaller == maybe {
		panic("couldn't decide")
	}
	return p.isLtSmaller == yes
}

func (p pair) checkIfLtIsSmaller(spaces string) pair {
	p = p.testForEmptySets()
	if p.isLtSmaller != maybe {
		return p
	}

	p = p.testForInts(spaces)
	if p.isLtSmaller != maybe {
		return p
	}

	if len(p.lt.subsets) > 0 && len(p.rt.subsets) > 0 {
		p = p.testFirstSubsets(spaces)
		if p.isLtSmaller != maybe {
			return p
		}

		p = p.testNextSubsets(spaces)
	}

	if p.isLtSmaller == maybe {
		p = p.isLessThan(len(p.lt.subsets), len(p.rt.subsets))
	}
	return p
}

func (p pair) testForEmptySets() pair {
	if !p.lt.isEmpty() && !p.rt.isEmpty() {
		return p
	}

	ltLength := p.lt.n + len(p.lt.subsets)
	rtLength := p.rt.n + len(p.rt.subsets)
	return p.isLessThan(ltLength, rtLength)
}

func (p pair) testForInts(spaces string) pair {
	if p.lt.isInt() && p.rt.isInt() {
		return p.isLessThan(p.lt.n, p.rt.n)
	}
	if p.lt.isInt() {
		newRt := p.rt.subsets[0].s
		return p.testSubsets(p.lt, newRt, spaces)
	}
	if p.rt.isInt() {
		newLt := p.lt.subsets[0].s
		return p.testSubsets(newLt, p.rt, spaces)
	}
	return p
}

func (p pair) testSubsets(lt set, rt set, spaces string) pair {
	newP := pair{lt: lt, rt: rt}
	newP = newP.checkIfLtIsSmaller(spaces + "  ")
	p.isLtSmaller = newP.isLtSmaller
	return p
}

func (p pair) testFirstSubsets(spaces string) pair {
	newLt := p.lt.subsets[0].s
	newRt := p.rt.subsets[0].s
	return p.testSubsets(newLt, newRt, spaces)
}

func (p pair) testNextSubsets(spaces string) pair {
	newLt := set{}.defaultSet()
	newLt.subsets = p.lt.subsets[1:]
	if len(newLt.subsets) == 1 {
		newLt = newLt.subsets[0].s
	}

	newRt := set{}.defaultSet()
	newRt.subsets = p.rt.subsets[1:]
	if len(newRt.subsets) == 1 {
		newRt = newRt.subsets[0].s
	}

	return p.testSubsets(newLt, newRt, spaces)
}

func (p pair) isLessThan(lt int, rt int) pair {
	switch true {
	case lt < rt:
		p.isLtSmaller = yes
	case lt > rt:
		p.isLtSmaller = no
	default:
		p.isLtSmaller = maybe
	}
	return p
}
