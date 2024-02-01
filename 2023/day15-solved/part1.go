package day15

import "strings"

func runA(input string) int {
	sum := 0
	steps := strings.Split(input, ",")
	for _, step := range steps {
		sum += hashString(step)
	}
	return sum
}

func hashString(str string) int {
	hash := 0
	for _, char := range str {
		hash += int(char)
		hash *= 17
		hash = hash % 256
	}
	return hash
}
