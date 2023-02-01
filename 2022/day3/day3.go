package day3

func SumPrioritiesOfBadges(sacks []string) int {
	sum := 0
	for i := 0; i < len(sacks); i = i + 3 {
		sackA := sacks[i]
		sackB := sacks[i+1]
		sackC := sacks[i+2]
		same := ""
		for _, a := range sackA {
			for _, b := range sackB {
				for _, c := range sackC {
					if a == b && b == c {
						same = string(a)
						break
					}
				}
			}
		}
		priority := priorities[same]
		sum = sum + priority
	}

	return sum
}

func SumPrioritiesOfMatchingTypes(sacks []string) int {
	sum := 0
	for _, v := range sacks {
		half := len(v) / 2
		sackA := v[0:half]
		sackB := v[half:len(v)]
		var same = ""
		for _, a := range sackA {
			for _, b := range sackB {
				if a == b {
					same = string(a)
					break
				}
			}
		}
		priority := priorities[same]
		sum = sum + priority
	}

	return sum
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
