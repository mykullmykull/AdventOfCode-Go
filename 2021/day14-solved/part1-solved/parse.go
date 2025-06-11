package day

import "strings"

func parse(input []string) (map[string]int, map[string][2]string) {
	pairCount := make(map[string]int)
	evolutionRules := make(map[string][2]string)

	// Read the initial polymer template
	template := input[0]

	// Read the evolution rules
	for _, line := range input[2:] {
		parts := strings.Split(line, " -> ")
		pair := parts[0]
		insertion := parts[1]
		evolutionRules[pair] = [2]string{pair[:1] + insertion, insertion + pair[1:]}
	}

	// Initialize the pair count from the template
	for pair := range evolutionRules {
		for x := 0; x < len(template)-1; x++ {
			if template[x:x+2] == pair {
				pairCount[pair]++
			}
		}
	}

	return pairCount, evolutionRules
}
