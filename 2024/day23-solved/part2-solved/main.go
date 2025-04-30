package day

import (
	"fmt"
	"goAoc2024/utils"
)

func main(input []string) string {
	network, all := parseNetwork(input)
	maxGroupLength := 0
	maxGroup := []string{}
	indexes := []int{0}
	for len(indexes) > 0 {
		if hasAllUnique(indexes) && hasAllPairs(network, all, indexes) {
			maxGroupLength = utils.UpdateMost(maxGroupLength, len(indexes))
			if maxGroupLength == len(indexes) {
				maxGroup = []string{}
				for _, x := range indexes {
					maxGroup = append(maxGroup, all[x])
				}
			}
			indexes = increment(append(indexes, indexes[len(indexes)-1]), len(all))
			continue
		}
		indexes = increment(indexes, len(all))
	}
	return alphabetize(maxGroup)
}

func hasAllUnique(indexes []int) bool {
	m := map[int]bool{}
	for _, x := range indexes {
		if m[x] {
			return false
		}
		m[x] = true
	}
	return true
}

func hasAllPairs(network map[string]bool, all []string, indexes []int) bool {
	for x := range indexes {
		c1 := all[indexes[x]]
		for y := x + 1; y < len(indexes); y++ {
			c2 := all[indexes[y]]
			p1 := fmt.Sprintf("%s-%s", c1, c2)
			p2 := fmt.Sprintf("%s-%s", c2, c1)
			_, ok1 := network[p1]
			_, ok2 := network[p2]
			if !ok1 && !ok2 {
				return false
			}
		}
	}
	return true
}

func increment(indexes []int, max int) []int {
	if len(indexes) == 0 {
		return indexes
	}
	lastX := len(indexes) - 1
	newLast := indexes[lastX] + 1
	indexes = indexes[:lastX]
	if newLast >= max {
		return increment(indexes, max)
	}
	return append(indexes, newLast)
}

func alphabetize(group []string) string {
	sorted := []string{group[0]}
	for _, s1 := range group[1:] {
		inserted := false
		for i, s2 := range sorted {
			if s1 < s2 {
				sorted = append(sorted[:i], append([]string{s1}, sorted[i:]...)...)
				inserted = true
				break
			}
		}
		if inserted {
			continue
		}
		sorted = append(sorted, s1)
	}
	str := utils.JoinArray(sorted, ",")
	return str[:len(str)-1]
}
