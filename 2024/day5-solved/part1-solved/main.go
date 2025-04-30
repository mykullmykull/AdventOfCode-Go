package day

import (
	"goAoc2024/utils"
	"strings"
)

type rule struct {
	bef, aft int
}

type update struct {
	str        string
	pagesOrder map[int]int
}

func main(input []string) int {
	total := 0
	rules, updates := parse(input)
	for _, u := range updates {
		println()
		println(u.str)
		isValid := true
		for _, r := range rules {
			x, befFound := u.pagesOrder[r.bef]
			y, aftFound := u.pagesOrder[r.aft]
			if befFound && aftFound && x > y {
				isValid = false
				break
			}
		}
		if !isValid {
			continue
		}
		total += u.getMiddlePage()
	}
	return total
}

func (u update) getMiddlePage() int {
	index := len(u.pagesOrder) / 2
	split := strings.Split(u.str, ",")
	str := split[index]
	page := utils.Atoi(str)
	return page
}
