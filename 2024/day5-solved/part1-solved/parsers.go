package day

import (
	"goAoc2024/utils"
	"strings"
)

func parse(input []string) ([]rule, []update) {
	rules := []rule{}
	updates := []update{}

	x := 0
	for x < len(input) {
		line := input[x]
		if line == "" {
			x++
			break
		}
		rules = append(rules, parseRule(line))
		x++
	}
	for x < len(input) {
		line := input[x]
		updates = append(updates, parseUpdate(line))
		x++
	}
	return rules, updates
}

func parseRule(line string) rule {
	split := strings.Split(line, "|")
	bef := utils.Atoi(split[0])
	aft := utils.Atoi(split[1])
	return rule{bef, aft}
}

func parseUpdate(line string) update {
	u := update{
		str:        line,
		pagesOrder: map[int]int{},
	}
	split := strings.Split(line, ",")
	for x, page := range split {
		u.pagesOrder[utils.Atoi(page)] = x
	}
	return u
}
