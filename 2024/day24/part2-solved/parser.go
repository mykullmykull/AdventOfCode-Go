package day

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(input []string) machine {
	// initialize machine
	m := machine{
		wireIndexes: map[string]int{},
		gateIndexes: map[gate]int{},
		toReplace:   map[string]string{},
	}

	// map gates and wires
	wireToGate := map[string]gate{}
	wireNumber := 0

	// Build the binary values from x and y wires
	xStr := ""
	yStr := ""

	for _, line := range input {
		// Parse x and y wires of the form "xn: value"
		if strings.Contains(line, ": ") {

			// Reset wireNumber for y wires
			if line[0:3] == "y00" {
				wireNumber = 0
			}

			// Split wire name (0) and value (1)
			parts := strings.Split(line, ": ")

			// Update x wires
			if line[0] == 'x' {

				// As wire number increases, prepend the value to xStr
				xStr = parts[1] + xStr

				// Map the wire name to a solved gate for later use
				wireToGate[getWireName("x", wireNumber)] = gate{parts[1], "", ""}

				// Found x wire, so increment wireNumber and continue to next line
				wireNumber++
				continue
			}

			// Update y wires
			// As wire number increases, prepend the value to yStr
			yStr = parts[1] + yStr

			// Map the wire name to a solved gate for later use
			wireToGate[getWireName("y", wireNumber)] = gate{parts[1], "", ""}

			// Found y wire, so increment wireNumber and continue to next line
			wireNumber++
			continue
		}

		// Skip the empty line separating xy wires and gates
		if !strings.Contains(line, " -> ") {
			continue
		}

		// Count the number of z wires for later use
		if strings.Contains(line, "z") {
			m.zCount++
		}

		// Split the line into gate (0) and wire name (1)
		parts := strings.Split(line, " -> ")

		// Build the gate object from the gate string
		gateStr := parts[0]
		gateParts := strings.Split(gateStr, " ")
		gate := gate{gateParts[0], gateParts[1], gateParts[2]}

		// Map the wire name to the gate for later use
		wireToGate[parts[1]] = gate
	}

	// Add z gates to wires and gates slices and index maps
	for z := 0; z < m.zCount; z++ {

		// Get wire name and gate
		wire := getWireName("z", z)
		gate := wireToGate[wire]

		// Update slices and index maps
		m.wires = append(m.wires, wire)
		m.gates = append(m.gates, gate)
		m.wireIndexes[wire] = len(m.wires) - 1
		m.gateIndexes[gate] = len(m.gates) - 1
	}

	// Add non z gates to wires and gates slices and index maps
	for wire, gate := range wireToGate {
		if wire[0] == 'z' {
			continue
		}

		// Update slices and index maps
		m.wires = append(m.wires, wire)
		m.gates = append(m.gates, gate)
		m.wireIndexes[wire] = len(m.wires) - 1
		m.gateIndexes[gate] = len(m.gates) - 1
	}

	// Caclulate z from x + y
	x64, xErr := strconv.ParseInt(xStr, 2, 64)
	y64, yErr := strconv.ParseInt(yStr, 2, 64)
	m.z = strconv.FormatInt(x64+y64, 2)
	if xErr != nil || yErr != nil {
		panic("failed to parse xStr " + xStr + " or yStr " + yStr)
	}
	m.x = int(x64)
	m.y = int(y64)

	return m
}

// Wire name format is prefix + number, where prefix is "x", "y", or "z", and number is always 2 digits
func getWireName(prefix string, index int) string {
	name := fmt.Sprintf("%s%d", prefix, index)
	if len(name) < 3 {
		name = fmt.Sprintf("%s0%d", prefix, index)
	}
	return name
}
