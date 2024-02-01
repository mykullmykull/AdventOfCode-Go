package day12

import (
	"fmt"
	"goAoc2023/utils"
	"strings"
)

func runA(input []string) int {
	sum := 0
	springs, groups := parseInput(input)
	for i := 0; i < len(springs); i++ {
		sum += countSolutions(springs[i], groups[i])
	}
	return sum
}

func parseInput(input []string) ([]string, [][]int) {
	springStrings := make([]string, len(input))
	groupsSlices := make([][]int, len(input))
	for i, line := range input {
		split := strings.Split(line, " ")
		springStrings[i] = split[0]
		groupStrsSplit := strings.Split(split[1], ",")
		groups := make([]int, len(groupStrsSplit))
		for j, groupStr := range groupStrsSplit {
			groups[j] = utils.Atoi(groupStr)
		}
		groupsSlices[i] = groups
	}
	return springStrings, groupsSlices
}

func countSolutions(springs string, groups []int) int {
	debug := false
	showMatches := true
	solutions := 0
	if showMatches {
		fmt.Printf("\n%v, %v\n", springs, groups)
	}
	spaces := make([]int, len(groups))
	for {
		if spaces[0] == -1 {
			break
		}
		if debug {
			fmt.Printf("  spaces: %v\n", spaces)
		}
		maxSpaces := getMaxSpaces(springs, groups)
		if noSpacesBetween(spaces) ||
			wrongNumberOfSpaces(len(spaces), len(groups)) ||
			tooLong(spaces, groups, len(springs)) {
			if debug {
				fmt.Println("    continuing")
			}

			spaces = getNextSpaces(spaces, maxSpaces)
			continue
		}

		if matches(springs, spaces, groups) {
			if showMatches {
				fmt.Printf("    spaces: %v matched!!\n", spaces)
			}
			solutions++
		}
		spaces = getNextSpaces(spaces, maxSpaces)
	}
	return solutions
}

func getMaxSpaces(springs string, groups []int) int {
	maxSpaces := len(springs)
	for _, g := range groups {
		maxSpaces -= g
	}
	maxSpaces -= len(groups) - 2
	return maxSpaces
}

func getNextSpaces(spaces []int, maxSpaces int) []int {
	debug := false
	indexToIncrease := 0
	for {
		if debug {
			fmt.Printf("index to increase: %d, value: %d\n", indexToIncrease, spaces[indexToIncrease])
		}
		if indexToIncrease >= len(spaces) {
			spaces[0] = -1
			if debug {
				fmt.Printf("final spaces: %v\n", spaces)
			}
			return spaces
		}
		if spaces[indexToIncrease]+1 > maxSpaces {
			spaces[indexToIncrease] = 0
			indexToIncrease++
			continue
		}
		spaces[indexToIncrease] += 1
		return spaces
	}
}

func noSpacesBetween(spacesSlice []int) bool {
	for _, n := range spacesSlice[1:] {
		if n == 0 {
			// fmt.Println("  no spaces between")
			return true
		}
	}
	return false
}

func wrongNumberOfSpaces(lenSpaces int, lenGroups int) bool {
	// if lenSpaces != lenGroups {
	// 	fmt.Printf("  wrong number of spaces, %d != %d\n", lenSpaces, lenGroups)
	// }
	return lenSpaces != lenGroups
}

func tooLong(spaces []int, groups []int, maxLength int) bool {
	total := 0
	for i := 0; i < len(spaces); i++ {
		total += spaces[i]
		total += groups[i]
	}
	// if total > maxLength {
	// 	fmt.Printf("  too long; total: %d, maxLength: %d\n", total, maxLength)
	// }
	return total > maxLength
}

func matches(springs string, spaces []int, groups []int) bool {
	debug := false
	if debug {
		fmt.Printf("  springs: %s, spaces: %v, groups: %v\n", springs, spaces, groups)
	}
	springIndex := 0
	for i := 0; i < len(spaces); i++ {
		if debug {
			fmt.Printf("    matching %d spaces\n", spaces[i])
		}
		for j := 0; j < spaces[i]; j++ {
			if springs[springIndex] == '#' {
				if debug {
					fmt.Println("    false, no space here")
				}
				return false
			}
			springIndex++
			if debug {
				fmt.Printf("      after space: %s\n", springs[:springIndex])
			}
		}
		if debug {
			fmt.Printf("    matching %d broken springs\n", groups[i])
		}
		for j := 0; j < groups[i]; j++ {
			if springs[springIndex] == '.' {
				if debug {
					fmt.Println("    false, no broken spring here")
				}
				return false
			}
			springIndex++
			if debug {
				fmt.Printf("      after broken spring: %s\n", springs[:springIndex])
			}
		}
	}
	for _, c := range springs[springIndex:] {
		if c == '#' {
			return false
		}
	}
	return true
}
