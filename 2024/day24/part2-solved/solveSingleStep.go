package day

import "strings"

var toSimplify = map[string]string{
	"0 AND 0": "0",
	"0 AND 1": "0",
	"1 AND 0": "0",
	"1 AND 1": "1",
	"0 ORR 0": "0",
	"0 ORR 1": "1",
	"1 ORR 0": "1",
	"1 ORR 1": "1",
	"0 XOR 0": "0",
	"0 XOR 1": "1",
	"1 XOR 0": "1",
	"1 XOR 1": "0",
}

func (m *machine) solveSingleStep(number int) {
	gateStrings := m.getGatesString()
	replaced := map[string]bool{}

	m.markXYToReplace(number)
	m.markSolvedToReplace()

	// Replace all solvable wires downstream
	originalLength := 0
	for len(gateStrings) != originalLength {
		originalLength = len(gateStrings)

		// Replace newly solved wires
		for k, v := range m.toReplace {
			gateStrings = strings.ReplaceAll(gateStrings, k, v)
			replaced[k] = true
		}
		m.toReplace = map[string]string{}

		// Simplify gates
		for k, v := range toSimplify {
			gateStrings = strings.ReplaceAll(gateStrings, k, v)
		}

		// Update newly solved wires
		gates := strings.Split(gateStrings, ",")
		for index, wire := range m.wires {
			gateParts := strings.Split(gates[index], " ")
			if replaced[wire] {
				continue
			}
			if len(gateParts) > 1 {
				continue
			}
			m.gates[index] = gate{gateParts[0], "", ""}
			if wire[0] == 'x' || wire[0] == 'y' || wire[0] == 'z' {
				continue
			}
			m.toReplace[wire] = gates[index]
		}
	}

}

func (m machine) getGatesString() string {
	gatesStrings := []string{}
	for _, g := range m.gates {
		gatesStrings = append(gatesStrings, g.toString())
	}
	return strings.Join(gatesStrings, ",")
}

func (m *machine) markXYToReplace(number int) {
	xWire := getWireName("x", number)
	xIndex := m.wireIndexes[xWire]
	xValue := m.gates[xIndex].toString()

	yWire := getWireName("y", number)
	yINdex := m.wireIndexes[yWire]
	yValue := m.gates[yINdex].toString()

	m.toReplace = map[string]string{
		xWire: xValue,
		yWire: yValue,
	}
}

func (m *machine) markSolvedToReplace() {
	// Also replace any intermediate wires that are already solved
	for index, wire := range m.wires {
		gateStr := m.gates[index].toString()
		if wire[0] == 'x' || wire[0] == 'y' || wire[0] == 'z' || len(gateStr) > 1 {
			continue
		}
		m.toReplace[wire] = gateStr
	}
}
