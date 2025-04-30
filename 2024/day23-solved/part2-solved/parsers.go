package day

import "strings"

func parseNetwork(input []string) (map[string]bool, []string) {
	network := map[string]bool{}
	computersMap := map[string]bool{}
	for _, line := range input {
		network[line] = true
		split := strings.Split(line, "-")
		computersMap[split[0]] = true
		computersMap[split[1]] = true
	}

	computers := []string{}
	for k := range computersMap {
		computers = append(computers, k)
	}
	return network, computers
}
