package day1

import (
	"sort"
	"strconv"
)

func TopThreeCalories(calStrs []string) int {
	calStrs = append(calStrs, "")
	top3 := []int{0, 0, 0}
	calories := 0
	for _, v := range calStrs {
		if v == "" {
			min := top3[0]
			if calories > min {
				top3[0] = calories
				sort.Ints(top3)
			}
			calories = 0
		} else {
			vint, _ := strconv.Atoi(v)
			calories = calories + vint
		}
	}
	return top3[0] + top3[1] + top3[2]
}

func MaxCalories(calStrs []string) int {
	max := 0
	calories := 0
	for _, v := range calStrs {
		if v == "" {
			if calories > max {
				max = calories
			}
			calories = 0
		} else {
			vint, _ := strconv.Atoi(v)
			calories = calories + vint
		}
	}
	return max
}
