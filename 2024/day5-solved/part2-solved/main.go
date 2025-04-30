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
		isValid := true
		for _, r := range rules {
			x, befFound := u.pagesOrder[r.bef]
			y, aftFound := u.pagesOrder[r.aft]
			if befFound && aftFound && x > y {
				isValid = false
				break
			}
		}
		if isValid {
			continue
		}
		u = u.sortStr(rules)
		middle := u.getMiddlePage()
		total += middle
	}
	return total
}

func (u update) sortStr(rules []rule) update {
	r := 0
	for r < len(rules) {
		rule := rules[r]
		x, befFound := u.pagesOrder[rule.bef]
		y, aftFound := u.pagesOrder[rule.aft]
		if befFound && aftFound && x > y {
			u = u.moveXBeforeY(x, y)
			r = 0
			continue
		}
		r++
	}
	return u
}

func (u update) getMiddlePage() int {
	index := len(u.pagesOrder) / 2
	split := strings.Split(u.str, ",")
	str := split[index]
	page := utils.Atoi(str)
	return page
}

func (u update) moveXBeforeY(x, y int) update {
	split := strings.Split(u.str, ",")
	newSplit := []string{}
	for z := 0; z < len(split); z++ {
		zStr := split[z]
		if z == x {
			continue
		}
		if z == y {
			newSplit = append(newSplit, split[x])
		}
		newSplit = append(newSplit, zStr)
	}
	u.str = strings.Join(newSplit, ",")
	for z, zStr := range newSplit {
		u.pagesOrder[utils.Atoi(zStr)] = z
	}
	return u
}
