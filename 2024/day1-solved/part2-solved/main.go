package day

import (
	"goAoc2024/utils"
)

func main(input []string) int {
	listA, listB := parse(input)
	total := 0
	for x := 0; x < len(listA); x++ {
		a := listA[x]
		count := count(listB, a)
		total += a * count
	}
	return total
}

func parse(input []string) ([]int, []int) {
	listA := []int{}
	listB := []int{}
	for _, line := range input {
		ints := utils.SplitInts(line, " ")
		listA = append(listA, ints[0])
		listB = append(listB, ints[1])
	}
	return listA, listB
}

func count(list []int, value int) int {
	count := 0
	for _, v := range list {
		if v == value {
			count++
		}
	}
	return count
}
