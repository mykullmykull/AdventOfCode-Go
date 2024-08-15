package day

import "fmt"

func part2(input []string) int {
	sumOfBadgePriorities := 0
	groupIndex := 0
	for groupIndex <= len(input)-3 {
		sumOfBadgePriorities += getBadgePriority(groupIndex, input)
		groupIndex += 3
	}
	return sumOfBadgePriorities
}

func getBadgePriority(groupIndex int, input []string) int {
	elf1 := input[groupIndex]
	elf2 := input[groupIndex+1]
	elf3 := input[groupIndex+2]
	for _, item1 := range elf1 {
		matchedItem2 := false
		for _, item2 := range elf2 {
			if item1 == item2 {
				matchedItem2 = true
				break
			}
		}
		if !matchedItem2 {
			continue
		}
		for _, item3 := range elf3 {
			if item1 == item3 {
				return priorities[string(item1)]
			}
		}
	}
	panic(fmt.Sprintf("no badge found between %s, %s, and %s", elf1, elf2, elf3))
}

var priorities = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}
