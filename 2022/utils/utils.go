package utils

import (
	"fmt"
	"math"
	"strconv"
)

func StrToInt(str string) int {
	out, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to an int", str))
	}
	return out
}

func UpdateMost(prev int, new int) int {
	if new > prev {
		return new
	}
	return prev
}

func UpdateLeast(prev int, new int) int {
	if new < prev {
		return new
	}
	return prev
}

func AbsForInt(n int) int {
	return int(math.Abs(float64(n)))
}
