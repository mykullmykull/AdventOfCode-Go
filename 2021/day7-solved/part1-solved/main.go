package day

import (
	"goAoc2021/utils"
	"math"
)

func main(input []int) int {
	min, max := getMinMax(input)
	minCost := math.MaxInt
	minPos := min
	for x := min; x <= max; x++ {
		cost := calculateCost(input, x)
		println("x:", x, "cost:", cost)
		if cost < minCost {
			minCost = cost
			minPos = x
		}
	}
	println("best x:", minPos)
	return minCost
}

func getMinMax(input []int) (int, int) {
	min := math.MaxInt
	max := 0
	for _, x := range input {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max
}

func calculateCost(input []int, pos int) int {
	cost := 0
	for _, x := range input {
		cost += utils.Abs(x - pos)
	}
	return cost
}
