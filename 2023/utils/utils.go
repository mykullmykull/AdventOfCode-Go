package utils

import (
	"fmt"
	"strconv"
)

func Atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Could not parse %s to int", str))
	}
	return i
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
