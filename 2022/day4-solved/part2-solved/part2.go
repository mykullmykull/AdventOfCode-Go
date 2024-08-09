package day

import (
	"fmt"
	"goAoc2022/utils"
	"strings"
)

func part2(input []string) int {
	completeOverlaps := 0
	for _, line := range input {
		if overlapsAtAll(line) {
			completeOverlaps++
		}
	}
	return completeOverlaps
}

func overlapsAtAll(line string) bool {
	elvesSplit := strings.Split(line, ",")
	elf1Split := strings.Split(elvesSplit[0], "-")
	elf2Split := strings.Split(elvesSplit[1], "-")
	min1 := utils.StrToInt(elf1Split[0])
	max1 := utils.StrToInt(elf1Split[1])
	min2 := utils.StrToInt(elf2Split[0])
	max2 := utils.StrToInt(elf2Split[1])
	overlaps := (min1 >= min2 && min1 <= max2) ||
		(max1 >= min2 && max1 <= max2) ||
		(min2 >= min1 && min2 <= max1) ||
		(max2 >= min1 && max2 <= max1)
	fmt.Printf("min1: %d, max1: %d, min2 %d, max2: %d, overlaps: %v\n", min1, max1, min2, max2, overlaps)
	return overlaps
}
