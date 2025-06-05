package day

import (
	"fmt"
	"slices"
)

type point struct {
	row int
	col int
}

func main(input []string) int {
	lowPoints := getLowPoints(input)
	basinSizes := make([]int, len(lowPoints))
	for i, lowPoint := range lowPoints {
		basinSizes[i] = getBasinSize(input, lowPoint)
	}
	slices.Sort(basinSizes)
	slices.Reverse(basinSizes)
	fmt.Println("Basin sizes:", basinSizes)
	basinProduct := 1
	for _, size := range basinSizes[:3] {
		basinProduct *= size
	}
	return basinProduct
}
