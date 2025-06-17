package day

import (
	"fmt"
	"goAoc2021/utils"
	"strings"
)

func explode(s string) string {
	// find first pair of numbers at depth 5
	depth := 0
	for x := range len(s) {
		if s[x] == '[' {
			depth++
		}
		if s[x] == ']' {
			depth--
		}
		if depth >= 5 && isDigit(s[x]) && isExplodable(s, x) {
			return explodePair(s, x-1)
		}
	}
	return s
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isExplodable(s string, start int) bool {
	foundComma := false
	foundClose := false
	for start < len(s) {
		if s[start] == ']' {
			foundClose = true
			return foundComma && foundClose
		}
		if s[start] == ',' {
			foundComma = true
		}
		if s[start] == '[' {
			// we found a new pair, so we can't explode this one
			return false
		}
		start++
	}
	return false
}

func explodePair(s string, start int) string {
	leftStr := s[:start]
	pair, rightStr := splitPairFromRight(s[start+1:])
	split := strings.Split(pair, ",")
	leftVal := utils.Atoi(split[0])
	rightVal := utils.Atoi(split[1])
	leftStr = addToLeft(leftStr, leftVal)
	rightStr = addToRight(rightStr, rightVal)
	return leftStr + "0" + rightStr
}

func splitPairFromRight(s string) (string, string) {
	pair := ""
	for x, b := range s {
		if b == ']' {
			pair += "]"
			return pair[:len(pair)-1], s[x+1:]
		}
		if b == '[' {
			panic("unexpected '[' in explodePair " + s)
		}
		pair += string(b)
	}
	panic("could not find closing ']' in explodePair " + s)
}

func addToLeft(s string, val int) string {
	rightStr := ""
	num := -1
	for y := range len(s) {
		x := len(s) - 1 - y
		char := s[x]
		charStr := s[x : x+1]
		if isDigit(char) {
			toAdd := utils.Atoi(charStr)
			if num >= 0 {
				num += toAdd * 10
			} else {
				num = toAdd
			}
			continue
		}
		if num >= 0 {
			return fmt.Sprintf("%s%d%s", s[:x+1], num+val, rightStr)
		}
		rightStr = charStr + rightStr
	}
	return s
}

func addToRight(s string, val int) string {
	leftStr := ""
	num := -1
	for x, char := range s {
		charStr := s[x : x+1]
		if isDigit(byte(char)) {
			toAdd := utils.Atoi(charStr)
			if num >= 0 {
				num *= 10
				num += toAdd
			} else {
				num = toAdd
			}
			continue
		}
		if num >= 0 {
			return fmt.Sprintf("%s%d%s", leftStr, num+val, s[x:])
		}
		leftStr = leftStr + charStr
	}
	return s
}
