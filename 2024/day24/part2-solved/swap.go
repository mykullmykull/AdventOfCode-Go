package day

func (m *machine) swap(a, b string) {
	m.wires[m.wireIndexes[a]], m.wires[m.wireIndexes[b]] = m.wires[m.wireIndexes[b]], m.wires[m.wireIndexes[a]]
	m.wireIndexes[a], m.wireIndexes[b] = m.wireIndexes[b], m.wireIndexes[a]
	m.swaps = append(m.swaps, a, b)
}
