package day

import (
	"fmt"
)

func main(input []string) int {
	wireValues, gates, unknowns := parse(input)
	for len(unknowns) > 0 {
		u := unknowns[0]
		unknowns = unknowns[1:]
		g := gates[u]
		value := g.getValue(wireValues)
		if value == -1 {
			unknowns = append(unknowns, u)
			continue
		}
		wireValues[u] = value
	}

	// for w, v := range wireValues {
	// 	println(w, v)
	// }

	return getBinaryFromZs(wireValues)
}

func (g gate) getValue(wireValues map[string]int) int {
	v1 := wireValues[g.w1]
	v2 := wireValues[g.w2]
	if v1 == -1 || v2 == -1 {
		return -1
	}
	switch g.op {
	case AND:
		return v1 & v2
	case OR:
		return v1 | v2
	case XOR:
		return v1 ^ v2
	default:
		panic(fmt.Sprintf("Unknown op: %d", g.op))
	}
}

func getBinaryFromZs(wireValues map[string]int) int {
	zWires := []string{"z00"}
	values := []int{wireValues["z00"]}
	for w, v := range wireValues {
		// fmt.Printf("zWires: %v\n", zWires)
		if w[0] != 'z' {
			continue
		}

		if w == "z00" {
			continue
		}

		inserted := false
		for x, w2 := range zWires {
			// println("  w", w, "w2", w2, w < w2)
			if w < w2 {
				zWires = append(zWires[:x], append([]string{w}, zWires[x:]...)...)
				values = append(values[:x], append([]int{v}, values[x:]...)...)
				inserted = true
				break
			}
		}
		if inserted {
			continue
		}
		zWires = append(zWires, w)
		values = append(values, v)
	}

	// fmt.Printf("zWires: %v\n", zWires)
	// fmt.Printf("values: %v\n", values)
	bin := 0
	for x, v := range values {
		bin += v << x
	}
	return bin
}
