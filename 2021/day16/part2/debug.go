package day

func (p *packet) printPacket() {
	// id              int
	// value           int
	// parent          *packet
	// operationType   int
	// childLengthType int
	// childLastIndex  int
	// childCount      int
	// children        []*packet
	println("Packet ID:", p.id)
	p.printParent()
	println("  Version:", p.version)
	p.printOperationType()

	if p.operationType == literal {
		println("  Literal Value:", p.value)
	}

	if p.childLengthType == bitLength {
		println("  Bit Length Type, last index:", p.childLastIndex)
		print("  Children:")
		p.printChildren()
	}

	if p.childLengthType == countLength {
		println("  Count Length Type, child count:", p.childCount)
		p.printChildren()
	}
}

func (p *packet) printParent() {
	if p.parent != nil {
		println("  Parent ID:", p.parent.id)
	} else {
		println("  No Parent")
	}
}

func (p *packet) printOperationType() {
	switch p.operationType {
	case sum:
		println("  Operation Type: Sum")
	case product:
		println("  Operation Type: Product")
	case minimum:
		println("  Operation Type: Minimum")
	case maximum:
		println("  Operation Type: Maximum")
	case literal:
		println("  Operation Type: Literal")
	case greaterThan:
		println("  Operation Type: Greater Than")
	case lessThan:
		println("  Operation Type: Less Than")
	case equalTo:
		println("  Operation Type: Equal To")
	default:
		println("  Unknown Operation Type:", p.operationType)
	}
}

func (p *packet) printChildren() {
	for _, child := range p.children {
		print(" ", child.id)
	}
	println()
}
