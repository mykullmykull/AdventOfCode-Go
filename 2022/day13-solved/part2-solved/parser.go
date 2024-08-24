package day

import (
	"fmt"
	"strconv"
)

func parse(input []string) []set {
	sets := []set{}
	for _, line := range input {
		if line == "" {
			continue
		}
		sets = append(sets, parseSet(line))
	}
	return sets
}

func parseSet(str string) set {
	s := set{}.defaultSet()

	// remove first [ and last ]
	str = trimOuterBrackets(str)

	// maybe empty
	if str == "" {
		return s
	}

	// maybe single int
	n, err := strconv.Atoi(str)
	if err == nil {
		s.n = n
		return s
	}

	// we know it has subsets
	newSubsets := []subset{}
	newSet := set{}
	for len(str) > 0 {
		str, newSet = extractNextSet(str)
		newSubsets = append(newSubsets, subset{newSet})
	}
	s.subsets = newSubsets
	return s
}

func trimOuterBrackets(str string) string {
	if str == "" {
		return str
	}

	if str[0] == '[' {
		str = str[1:]
	}

	last := len(str) - 1
	if str[last] == ']' {
		str = str[:last]
	}
	return str
}

func extractNextSet(str string) (string, set) {
	nextSetStr := ""
	brackets := 0
	for len(str) > 0 {
		char := str[0]
		// println("    brackets", brackets, "nextSetStr", nextSetStr, "str", str)
		if brackets == 0 && char == ',' {
			return str[1:], parseSet(nextSetStr)
		}

		if char == '[' {
			brackets++
		}
		if char == ']' {
			brackets--
		}
		nextSetStr += string(char)
		str = str[1:]
	}
	return str, parseSet(nextSetStr)
}

func setsToString(sets []set) string {
	str := ""
	for _, s := range sets {
		str += fmt.Sprintf("%v\n", s)
	}
	return str
}
