package day

import "fmt"

func (m *monitor) printState() {
	println()
	m.printDigits()
	println()
	m.printSegs()
	println()
	println("\n==========================\n")
}

func (m monitor) printDigits() {
	for index, d := range m.digits {
		if len(d) == 0 {
			continue
		}
		fmt.Printf("wStr: %s, Possible Digits: %v\n", m.wStrs[index], d)
	}
}

func (m monitor) printSegs() {
	wires := allWires()
	for index, s := range m.segs {
		// if len(s) == 0 {
		// 	continue
		// }
		fmt.Printf("Wire: %c, Possible Segs: %s\n", wires[index], s)
	}
}
