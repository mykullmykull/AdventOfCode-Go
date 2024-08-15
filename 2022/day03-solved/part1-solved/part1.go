package day

import (
	"fmt"
)

func part1(input []string) int {
	repeatedPrioritySum := 0
	for _, line := range input {
		repeatedItem := findRepeatedChar(line)
		repeatedPrioritySum += priorities[repeatedItem]
	}
	return repeatedPrioritySum
}

func findRepeatedChar(line string) string {
	firstCompartment := line[:len(line)/2]
	secondCompartment := line[len(line)/2:]
	for _, c1 := range firstCompartment {
		for _, c2 := range secondCompartment {
			if c1 == c2 {
				return string(c1)
			}
		}
	}
	panic(fmt.Sprintf("no repeat found between %s and %s", firstCompartment, secondCompartment))
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
