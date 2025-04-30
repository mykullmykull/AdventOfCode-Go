package day

import (
	"goAoc2024/utils"
)

const (
	dsc   = -1
	unset = 0
	asc   = 1
)

func main(input []string) int {
	reports := parse(input)
	safeCount := 0
	for _, report := range reports {
		for x := -1; x < len(report); x++ {
			isSafe := isSafeWithoutLevel(report, x)
			if isSafe {
				safeCount++
				break
			}
		}
	}
	return safeCount
}

func parse(input []string) [][]int {
	reports := [][]int{}
	for _, line := range input {
		report := utils.SplitInts(line, " ")
		reports = append(reports, report)
	}
	return reports
}

func isSafeWithoutLevel(report []int, x int) bool {
	if x == -1 {
		return isSafe(report)
	}
	c := make([]int, len(report))
	copy(c, report)
	return isSafe(append(c[:x], c[x+1:]...))
}

func isSafe(report []int) bool {
	slope := unset
	for x := range report {
		// count if last element
		if x == len(report)-1 {
			return true
		}

		this := report[x]
		next := report[x+1]

		// set slope
		if this == next {
			return false
		}
		if x == 0 && this < next {
			slope = asc
		}
		if x == 0 && this > next {
			slope = dsc
		}

		// break if unsafe
		if slope == asc && this >= next {
			return false
		}
		if slope == dsc && this <= next {
			return false
		}
		if utils.Abs(this-next) > 3 {
			return false
		}
	}
	return true
}
