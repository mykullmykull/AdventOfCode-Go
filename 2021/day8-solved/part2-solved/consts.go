package day

import "slices"

func allDigitsStr() string {
	return "0123456789"
}

func allDigits() []string {
	return []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
}

func allLitSegs() []string {
	return []string{
		"ABCEFG",
		"CF",
		"ACDEG",
		"ACDFG",
		"BCDF",
		"ABDFG",
		"ABDEFG",
		"ACF",
		"ABCDEFG",
		"ABCDFG",
	}
}

func digitToLitSegs(d string) string {
	return allLitSegs()[slices.Index(allDigits(), d)]
}

func litSegsToDigit(s string) string {
	return allDigits()[slices.Index(allLitSegs(), s)]
}

// allLights returns an interable slice of all lights
func allSegs() []rune {
	return []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
}

func allSegsStr() string {
	return "ABCDEFG"
}

// allWires returns an interable slice of all wires
func allWires() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
}

func allWiresStr() string {
	return "abcdefg"
}
