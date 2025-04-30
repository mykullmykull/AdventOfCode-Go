package day

import (
	"goAoc2024/utils"
)

func main(input []string) int {
	listA, listB := parse(input)
	total := 0
	for x := 0; x < len(listA); x++ {
		a := listA[x]
		b := listB[x]
		total += utils.Abs(a - b)
	}
	return total
}

func parse(input []string) ([]int, []int) {
	listA := []int{}
	listB := []int{}
	for _, line := range input {
		ints := utils.SplitInts(line, " ")
		listA = insertSorted(listA, ints[0])
		listB = insertSorted(listB, ints[1])
	}
	return listA, listB
}

func insertSorted(list []int, value int) []int {
	if len(list) == 0 {
		return []int{value}
	}
	for i, v := range list {
		if value < v {
			return append(list[:i], append([]int{value}, list[i:]...)...)
		}
	}
	return append(list, value)
}
