package day

type stream struct {
	binary        string
	bitIndex      int
	outerPacket   *packet
	currentParent *packet
	nextId        int
	packets       map[int]*packet
}

func (s *stream) checkIfFinishedChildrenBits() {
	if s.currentParent == nil {
		return
	}

	if s.currentParent.operationType == literal {
		panic("Current parent should never be a literal packet")
	}

	if s.currentParent.childLengthType != bitLength {
		return
	}

	remainingBits := s.currentParent.childLastIndex - s.bitIndex
	if remainingBits >= 11 {
		// We have enough bits to read the next packet
		return
	}

	s.bitIndex = s.currentParent.childLastIndex + 1
	s.currentParent = s.currentParent.parent
}

func (s *stream) checkIfFinishedChildrenCount() {
	if s.currentParent == nil {
		return
	}

	if s.currentParent.operationType == literal {
		panic("Current parent should never be a literal packet")
	}

	if s.currentParent.childLengthType != countLength {
		return
	}

	if len(s.currentParent.children) < s.currentParent.childCount {
		return // Not enough packets parsed yet
	}

	s.currentParent = s.currentParent.parent
}

func (s *stream) initPacket() *packet {
	// -1 indicates not yet set
	p := &packet{
		id:              s.nextId,
		startingIndex:   s.bitIndex,
		parent:          s.currentParent,
		value:           -1,
		childLengthType: -1,
	}
	s.nextId++

	if s.currentParent == nil {
		s.outerPacket = p
		return p
	}

	s.currentParent.children = append(s.currentParent.children, p)
	return p
}
