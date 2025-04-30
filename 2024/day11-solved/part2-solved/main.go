package day

import (
	"fmt"
	"goAoc2024/utils"
)

type stone struct {
	age int
	val string
}

func main(values []string) int {
	count := 0
	stones := parseStones(values)
	cache := map[stone]int{}
	for _, s := range stones {
		countProcessed, cache := processStone(s, cache)
		count += countProcessed
		cache[s] = countProcessed
	}
	return count
}

func parseStones(values []string) []stone {
	stones := []stone{}
	for _, v := range values {
		stones = append(stones, stone{age: 0, val: v})
	}
	return stones
}

func processStone(s stone, cache map[stone]int) (int, map[stone]int) {
	count := 1
	for s.age < 75 {
		if count, ok := cache[s]; ok {
			return count, cache
		}

		s.age++
		if s.val == "0" {
			s.val = "1"
			continue
		}

		if len(s.val)%2 == 1 {
			s.val = fmt.Sprintf("%d", utils.Atoi(s.val)*2024)
			continue
		}

		return processWithChildren(s, cache)
	}
	return count, cache
}

func processWithChildren(s stone, cache map[stone]int) (int, map[stone]int) {
	half := len(s.val) / 2
	leftVal := s.val[:half]
	left := stone{age: s.age, val: leftVal}
	fromLeft, cache := processStone(left, cache)
	cache[left] = fromLeft

	rightVal := removeLeadingZeros(s.val[half:])
	right := stone{age: s.age, val: rightVal}
	fromRight, cache := processStone(right, cache)
	cache[right] = fromRight

	return fromLeft + fromRight, cache
}

func removeLeadingZeros(s string) string {
	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
	}
	if len(s) == 0 {
		return "0"
	}
	return s
}
