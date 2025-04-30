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
		slope := unset
		for x := range report {
			// count if last element
			if x == len(report)-1 {
				safeCount++
				break
			}

			this := report[x]
			next := report[x+1]

			// set slope
			if this == next {
				break
			}
			if x == 0 && this < next {
				slope = asc
			}
			if x == 0 && this > next {
				slope = dsc
			}

			// break if unsafe
			if slope == asc && this >= next {
				break
			}
			if slope == dsc && this <= next {
				break
			}
			if utils.Abs(this-next) > 3 {
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
