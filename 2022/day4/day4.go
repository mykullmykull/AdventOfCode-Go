package day4

import (
	"strconv"
	"strings"
)

func CountOverlaps(pairs []string, part int) int {
	count := 0
	for _, pair := range pairs {
		a, b, _ := strings.Cut(pair, ",")

		aMinStr, aMaxStr, _ := strings.Cut(a, "-")
		aMin, _ := strconv.Atoi(aMinStr)
		aMax, _ := strconv.Atoi(aMaxStr)

		bMinStr, bMaxStr, _ := strings.Cut(b, "-")
		bMin, _ := strconv.Atoi(bMinStr)
		bMax, _ := strconv.Atoi(bMaxStr)

		aEncompassesB := aMin <= bMin && aMax >= bMax
		bEncompassesA := bMin <= aMin && bMax >= aMax
		overlapps := aMin <= bMax && aMax >= bMin
		if part == 1 && (aEncompassesB || bEncompassesA) {
			count++
		}

		if part == 2 && overlapps {
			count++
		}
	}
	return count
}
