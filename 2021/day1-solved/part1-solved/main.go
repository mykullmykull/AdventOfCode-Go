package day

import "math"

func main(input []int) int {
	increases := 0
	last := math.MaxInt
	for _, depth := range input {
		if depth > last {
			increases++
		}
		last = depth
	}
	return increases
}
