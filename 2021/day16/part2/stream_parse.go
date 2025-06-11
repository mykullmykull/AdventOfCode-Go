package day

import (
	"fmt"
	"strings"
)

func (s *stream) parseNextPacket() {
	println()
	// println(s.binary[:s.bitIndex], s.binary[s.bitIndex:])
	println("Parsing packet", s.nextId, "at index", s.bitIndex)

	if s.hasOnlyFillerBinaryLeft() {
		s.bitIndex = len(s.binary) // Move to the end to avoid further parsing
		return
	}

	// create packet with initial values
	p := s.initPacket()

	// get version (ignore)
	p.version = s.readBits(3)

	// get type ID
	p.operationType = s.readBits(3)

	// parse literal value
	if p.operationType == literal {
		p.value = s.parseLiteralValue()
		p.printPacket()
		return
	}

	// handle operator packet
	println("~~making current parent", p.id)
	s.currentParent = p
	p.childLengthType = s.readBits(1)
	if p.childLengthType == bitLength {
		childrenLength := s.readBits(15)
		p.childLastIndex = childrenLength + s.bitIndex
	} else {
		p.childCount = s.readBits(11)
	}
	p.printPacket()
	s.packets[p.id] = p
}

func (s *stream) hasOnlyFillerBinaryLeft() bool {
	if strings.Contains(s.binary[s.bitIndex:], "1") {
		return false
	}
	countBitsRemaining := len(s.binary) - s.bitIndex
	if countBitsRemaining < 11 { // minimum for literal packet
		return true
	}
	if s.binary[s.bitIndex+3:s.bitIndex+7] != "1001" && // count length packet
		countBitsRemaining < 18 { // minimum for count length packet
		return true
	}
	if s.binary[s.bitIndex+3:s.bitIndex+7] != "1000" && // bit length packet
		countBitsRemaining < 23 { // minimum for bit length packet
		return true
	}
	return false
}

func (s *stream) readBits(length int) int {
	println("Reading", length, "bits from index", s.bitIndex)

	// verify that we have enough bits left to read
	if length <= 0 || length > len(s.binary)-s.bitIndex {
		s.bitIndex = len(s.binary) // Move to the end to avoid further parsing
		return -1
	}

	// read the bits from the binary string
	value := 0
	for x := 0; x < length; x++ {
		fmt.Printf("%v", s.binary[s.bitIndex])
		value = (value << 1) | int(s.binary[s.bitIndex]-'0')
		s.bitIndex++
	}
	println(" -> value:", value)

	return value
}

func (s *stream) parseLiteralValue() int {
	value := 0
	for {
		group := s.readBits(5)
		// Increase the value by shifting left and adding the last 4 bits
		value = (value << 4) | (group & 0x0F) // Get the last 4 bits

		// Break if that was the last group
		if group&0x10 == 0 { // Last group if the first bit is 0
			break
		}
	}
	return value
}
