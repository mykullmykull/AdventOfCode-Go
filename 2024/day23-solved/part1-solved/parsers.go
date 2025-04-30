package day

import "strings"

type pair struct {
	c1 string
	c2 string
}

func parseNetwork(input []string) (map[pair]bool, []string, []string) {
	network := map[pair]bool{}
	computersMap := map[string]bool{}
	for _, line := range input {
		split := strings.Split(line, "-")
		p1 := pair{c1: split[0], c2: split[1]}
		p2 := pair{c1: split[1], c2: split[0]}
		network[p1] = true
		network[p2] = true
		computersMap[split[0]] = true
		computersMap[split[1]] = true
	}

	computers := []string{}
	ts := []string{}
	for k := range computersMap {
		computers = append(computers, k)
		if k[0] == 't' {
			ts = append(ts, k)
		}
	}
	return network, computers, ts
}
