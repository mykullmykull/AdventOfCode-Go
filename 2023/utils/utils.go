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
