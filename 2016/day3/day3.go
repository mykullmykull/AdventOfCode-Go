package day3

import (
	"fmt"
)

func runB(input []string) int {
	countTriangles := 0
	sides := parseSidesByColumn(input)

	for i := 0; i < len(sides); i += 3 {
		a := sides[i]
		b := sides[i+1]
		c := sides[i+2]
		if a+b > c && a+c > b && b+c > a {
			countTriangles++
		}
	}
	return countTriangles
}

func parseSidesByColumn(input []string) []int {
	col1 := make([]int, len(input))
	col2 := make([]int, len(input))
	col3 := make([]int, len(input))
	for _, row := range input {
		var a, b, c int
		fmt.Sscanf(row, "%d %d %d", &a, &b, &c)
		col1 = append(col1, a)
		col2 = append(col2, b)
		col3 = append(col3, c)
	}
	return append(append(col1, col2...), col3...)
}

func runA(input []string) int {
	countTriangles := 0
	for _, row := range input {
		if isTriangle(row) {
			countTriangles++
		}
	}
	return countTriangles
}

func isTriangle(row string) bool {
	var a, b, c int
	fmt.Sscanf(row, "%d %d %d", &a, &b, &c)
	return a+b > c && a+c > b && b+c > a
}
