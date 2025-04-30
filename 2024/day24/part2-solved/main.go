package day

import (
	"slices"
	"strings"
)

type machine struct {
	x           int
	y           int
	z           string
	toReplace   map[string]string
	zCount      int
	wires       []string
	wireIndexes map[string]int
	gates       []gate
	gateIndexes map[gate]int
	swaps       []string
}

func main(inputs []string) string {
	m := parse(inputs)
	m.getSwaps()
	slices.Sort(m.swaps)
	return strings.Join(m.swaps, ",")
}
