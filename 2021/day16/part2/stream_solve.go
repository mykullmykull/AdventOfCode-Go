package day

import "fmt"

func (s *stream) solve() {
	// start with the outer packet
	s.currentParent = s.outerPacket

	// process until the outer packet has a value
	for {
		p := s.currentParent
		println("focusing on packet:", p.id, "value:", p.value, "children:", len(p.children))
		// check that we have the right number of children for the operation type
		if p.id == 11 {
			println("  packet id:", p.id, "has value", p.value, "and operation type", p.operationType, "with children count", len(p.children))
			println("  p.operationType >= greaterThan:", p.operationType >= greaterThan)
		}
		if p.operationType >= greaterThan && len(p.children) != 2 {
			panic(fmt.Sprintf("  packet id: %d should have exactly 2 children, but has %d instead", p.id, len(p.children)))
		}

		// check if all children have values, if not, set focus to the next unsolved child and continue
		childrenNeedSolving := false
		for _, child := range p.children {
			if child.value == -1 {
				childrenNeedSolving = true
				s.currentParent = child
				break
			}
		}
		if childrenNeedSolving {
			continue
		}

		// perform the operation based on the operation type
		switch p.operationType {
		case literal:
			panic("  literal packets should already have a value")
		case sum:
			p.value = getSum(p.children)
		case product:
			p.value = getProduct(p.children)
		case minimum:
			p.value = getMinimum(p.children)
		case maximum:
			p.value = getMaximum(p.children)
		case greaterThan:
			p.value = getGreaterThan(p.children)
		case lessThan:
			p.value = getLessThan(p.children)
		case equalTo:
			p.value = getEqualTo(p.children)
		default:
			panic(fmt.Sprintf("  unknown value type: %d", p.operationType))
		}

		if p.parent == nil {
			// this is the outer packet, we are done
			s.outerPacket = p
			println("Outer packet value:", s.outerPacket.value)
			return
		}

		// move focus back tp the parent packet
		s.currentParent = p.parent
	}
}

func getSum(children []*packet) int {
	sum := 0
	for _, child := range children {
		sum += child.value
	}
	return sum
}

func getProduct(children []*packet) int {
	product := 1
	for _, child := range children {
		product *= child.value
	}
	return product
}

func getMinimum(children []*packet) int {
	min := children[0].value
	for _, child := range children[1:] {
		if child.value < min {
			min = child.value
		}
	}
	return min
}

func getMaximum(children []*packet) int {
	max := children[0].value
	for _, child := range children[1:] {
		if child.value > max {
			max = child.value
		}
	}
	return max
}

func getGreaterThan(children []*packet) int {
	if children[0].value > children[1].value {
		return 1
	}
	return 0
}

func getLessThan(children []*packet) int {
	if children[0].value < children[1].value {
		return 1
	}
	return 0
}

func getEqualTo(children []*packet) int {
	if children[0].value == children[1].value {
		return 1
	}
	return 0
}
