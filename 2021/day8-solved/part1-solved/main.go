package day

import (
	"slices"
	"strings"
)

func main(input []string) int {
	count := 0
	for _, line := range input {
		parts := strings.Split(line, "|")
		data := strings.Split(parts[1], " ")
		for _, str := range data {
			if slices.Contains([]int{2, 3, 4, 7}, len(str)) {
				count++
			}
		}
	}
	return count
}
