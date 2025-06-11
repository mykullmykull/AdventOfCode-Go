package day

const (
	// Length Types
	bitLength   = 0
	countLength = 1

	// Operator Types
	sum         = 0
	product     = 1
	minimum     = 2
	maximum     = 3
	literal     = 4
	greaterThan = 5
	lessThan    = 6
	equalTo     = 7
)

type packet struct {
	version         int
	id              int
	startingIndex   int
	value           int
	parent          *packet
	operationType   int
	childLengthType int
	childLastIndex  int
	childCount      int
	children        []*packet
}

func main(hex string) int {
	s := &stream{
		binary:  convertHexToBin(hex),
		packets: make(map[int]*packet),
	}
	println("Binary:", s.binary)

	for len(s.binary)-s.bitIndex > 0 {
		s.checkIfFinishedChildrenBits()
		s.checkIfFinishedChildrenCount()
		s.parseNextPacket()
	}

	// s.packets[11].printPacket()

	s.solve()

	return s.outerPacket.value
}
