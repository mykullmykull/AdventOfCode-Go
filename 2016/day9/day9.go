package day9

import (
	"strings"

	"goAoc2016/utils"
)

func runB(input string) int {
	length := 0
	for i := 0; i < len(input); i++ {
		// fmt.Printf("~~~ remaining: %s, i: %d, length: %d\n", input[i:], i, length)
		if input[i] == '(' {
			toRepeat, multiplier, indexesToSkip := parseMarker(input[i:])
			// fmt.Printf("~~~ multiplier: %d, indexesToSkip: %d, toRepeat: %s\n", multiplier, indexesToSkip, toRepeat)
			length += runB(toRepeat) * multiplier
			i += indexesToSkip
			continue
		}
		length++
	}
	return length
}

func isFinished(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			return false
		}
	}
	return true
}

func runA(input string) int {
	decompressed := decompress(input)
	return len(decompressed)
}

func decompress(input string) string {
	decompressed := ""
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			toRepeat, repeatTimes, indexesToSkip := parseMarker(input[i:])
			for j := 0; j < repeatTimes; j++ {
				decompressed += toRepeat
			}
			i += indexesToSkip
			continue
		}
		decompressed += string(input[i])
	}
	return decompressed
}

func parseMarker(str string) (string, int, int) {
	marker := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ')' {
			marker += string(str[i])
			break
		}
		marker += string(str[i])
	}
	split := strings.Split(marker[1:len(marker)-1], "x")
	countToRepeat := utils.Atoi(split[0])
	toRepeat := str[len(marker) : len(marker)+countToRepeat]
	repeatTimes := utils.Atoi(split[1])
	indexesToSkip := len(marker) + len(toRepeat) - 1

	return toRepeat, repeatTimes, indexesToSkip
}
