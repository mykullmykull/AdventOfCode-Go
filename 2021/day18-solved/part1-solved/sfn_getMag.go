package day

import (
	"goAoc2021/utils"
	"strings"
)

func getMagnitude(s string) int {
	for strings.Contains(s, "[") {
		s = getNextMagnitude(s)
	}
	return utils.Atoi(s)
}

func getNextMagnitude(s string) string {
	end := strings.Index(s, "]")
	pair := getPair(s, end)
	mag := getMagnitudeOfPair(pair)
	s = strings.ReplaceAll(s, pair, utils.Itoa(mag))
	return s
}

func getPair(s string, end int) string {
	start := end - 1
	for start >= 0 {
		char := s[start]
		if char == ']' {
			start--
			continue
		}

		if s[start] == '[' {
			return s[start : end+1]
		}
		start--
	}
	panic("No pair found in string: " + s)
}

func getMagnitudeOfPair(s string) int {
	s = s[1 : len(s)-1] // remove outer brackets
	split := strings.Split(s, ",")
	ltVal := utils.Atoi(split[0])
	rtVal := utils.Atoi(split[1])
	return 3*ltVal + 2*rtVal
}
