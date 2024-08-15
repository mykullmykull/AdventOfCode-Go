package day

import (
	"goAoc2022/utils"
	"strings"
)

func part1(input []string) int {
	completeOverlaps := 0
	for _, line := range input {
		if completelyOverlaps(line) {
			completeOverlaps++
		}
	}
	return completeOverlaps
}

func completelyOverlaps(line string) bool {
	elvesSplit := strings.Split(line, ",")
	elf1Split := strings.Split(elvesSplit[0], "-")
	elf2Split := strings.Split(elvesSplit[1], "-")
	min1 := utils.StrToInt(elf1Split[0])
	max1 := utils.StrToInt(elf1Split[1])
	min2 := utils.StrToInt(elf2Split[0])
	max2 := utils.StrToInt(elf2Split[1])
	overlaps := (min1 <= min2 && max1 >= max2) || (min2 <= min1 && max2 >= max1)
	// fmt.Printf("\nmin1: %d, min2 %d, max1: %d, max2: %d, overlaps: %v", min1, min2, max1, max2, overlaps)
	return overlaps
}
