package day

import (
	"fmt"
)

func (m *machine) getSwaps() {
	// Store the value to carry to the next column
	nextC := ""
	for z := 0; z < m.zCount; z++ {
		wireName := getWireName("z", z)
		index := m.wireIndexes[wireName]
		currentGate := m.gates[index]
		expectedGate := m.getZnGate(z, nextC)

		// If the gate at the current index isn't the expected value, we need to swap
		if !currentGate.isFullMatch(expectedGate) {

			// If z wire needs to be swapped, this will return the other wire,
			// If an internal wire needs to be swapped for this z wire to work, then this will handle that swap and return the z wire
			wireToSwap := m.getWireAndMaybeSwapInternal(expectedGate)

			// If z wire was returned again, then the swap already happened inside getWire
			if wireName != wireToSwap {
				m.swap(wireName, wireToSwap)
				index = m.wireIndexes[wireName]
			}
		}

		// Double check that the gate resolves to the right value now, panic if not
		m.solveSingleStep(z)
		expectedValue := m.z[len(m.z)-z-1 : len(m.z)-z]
		if m.gates[index].toString() != expectedValue {
			panic(fmt.Sprintf("%s not solved: %s != %s", wireName, m.gates[index].toString(), expectedValue))
		}

		// Get the value to carry to the next column
		nextC = m.getCn(z, nextC)
	}
}

// getZnGate returns the expected gate for the z wire at index n
// zn = (xn XOR yn) XOR cn-1
func (m *machine) getZnGate(n int, c string) gate {
	// Final z wire won't have a defined gate, it will just be the last carry value
	if n == m.zCount-1 {
		g := m.gates[m.wireIndexes[c]]
		return g
	}

	// Get the x and y wires for the current index n as well as their XOR gate
	xWire := getWireName("x", n)
	yWire := getWireName("y", n)
	xorGate := gate{xWire, "XOR", yWire}

	// The 1st z wire has no carry value, so we just return the XOR gate
	if c == "" {
		return xorGate
	}

	// Find the wire name that encompasses the expected XOR gate
	xorWire := m.getWireAndMaybeSwapInternal(xorGate)
	return gate{xorWire, "XOR", c}
}

// getCn returns the expected carry value for the next column
// cn = (xn AND yn) ORR (cn-1 AND (xn XOR yn))
func (m *machine) getCn(n int, c string) string {
	// Final z wire won't have a defined carry value
	if n == m.zCount-1 {
		return "finished"
	}

	// Build the x and y wire names for the current index n as well as their AND gate
	xWire := getWireName("x", n)
	yWire := getWireName("y", n)
	andGate := gate{xWire, "AND", yWire}

	// Find the wire name that encompasses the expected AND gate
	andWire := m.getWireAndMaybeSwapInternal(andGate)

	// The 1st z wire has no carry value, so we just return the AND gate
	if c == "" {
		return andWire
	}

	// Build the expected XOR gate for the x and y wires, then find the wire name that encompasses it
	xorGate := gate{xWire, "XOR", yWire}
	xorWire := m.getWireAndMaybeSwapInternal(xorGate)

	// Build the expected AND gate for the c and xorWire, then find the wire name that encompasses it
	candGate := gate{c, "AND", xorWire}
	candWire := m.getWireAndMaybeSwapInternal(candGate)

	// Build the expected ORR gate for the andWire and candWire, then find and return the wire name that encompasses it
	orrGate := gate{andWire, "ORR", candWire}
	orrWire := m.getWireAndMaybeSwapInternal(orrGate)

	return orrWire
}

func (m *machine) getWireAndMaybeSwapInternal(g gate) string {
	// Return it if the gate needs no swap
	index, ok := m.gateIndexes[g]
	if ok {
		return m.wires[index]
	}

	// Wire position is ambiguous, try swapping the position
	g = gate{g.b, g.opp, g.a}
	index, ok = m.gateIndexes[g]
	if ok {
		return m.wires[index]
	}

	// Loop through all gates and look for a match on opp and at least one of wire a or b
	for _, gate := range m.gates {
		// See if the gate is a match
		swap1, swap2, err := gate.swapToMatch(g)
		if err != nil {
			continue
		}

		// If the gate is a match, swap the wires in the machine and the gate
		m.swap(swap1, swap2)
		g.swap(swap1, swap2)

		// Return the fixed wire after the swap
		index, ok = m.gateIndexes[g]
		if ok {
			return m.wires[index]
		}
	}

	panic(fmt.Sprintf("wire not found: %s", g.toString()))
}
