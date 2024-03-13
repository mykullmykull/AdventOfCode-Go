package day19

import (
	"fmt"
	"goAoc2023/utils"
	"strings"
)

func runA(input []string) int {
	debug := true
	piles, parts := parseInput(input)
	if debug {
		for k, v := range piles {
			log(fmt.Sprintf("name: %s, piles: %v", k, v.rules), debug)
		}
		for _, p := range parts {
			fmt.Println(p)
		}
	}
	totalAcceptedAttributes := 0
	log("", debug)
	for _, p := range parts {
		log(fmt.Sprintf("part: %v", p), debug)
		p = p.getFinalDestination(piles)
		if p.pileName == "A" {
			totalAcceptedAttributes += p.totalAttributes()
		}
	}
	return totalAcceptedAttributes
}

func parseInput(input []string) (map[string]pile, []part) {
	debug := false
	piles := map[string]pile{}
	parts := []part{}
	cutOff := -1
	for i, line := range input {
		log(fmt.Sprintf("\npile line: %s", line), debug)
		if line == "" {
			cutOff = i
			break
		}
		name, pile := parsePile(line)
		log(fmt.Sprintf("name: %s, pile: %v", name, pile), debug)
		piles[name] = pile
	}
	for i := cutOff + 1; i < len(input); i++ {
		log(fmt.Sprintf("\npart line: %s", input[i]), debug)
		part := parsePart(input[i])
		log(fmt.Sprintf("part: %v", part), debug)
		parts = append(parts, part)
	}

	return piles, parts
}

func parsePile(line string) (string, pile) {
	debug := false
	split := strings.Split(line, "{")
	name := split[0]
	rulesStr := strings.Split(strings.ReplaceAll(split[1], "}", ""), ",")
	log(fmt.Sprintf("name: %s, rulesStr: %v", name, rulesStr), debug)
	rules := []rule{}
	for _, r := range rulesStr {
		log(fmt.Sprintf("rule: %s", r), debug)
		newRule := parseRule(r)
		log(fmt.Sprintf("newRule: %v", newRule), debug)
		rules = append(rules, newRule)
	}
	return name, pile{rules: rules}
}

func parseRule(line string) rule {
	debug := false
	if !strings.Contains(line, ":") {
		return rule{destination: line}
	}

	attribute := line[:1]
	operator := line[1:2]
	split := strings.Split(line[2:], ":")
	value := utils.Atoi(split[0])
	destination := split[1]
	log(fmt.Sprintf("attribute: %s, operator: %s, split: %v, value: %d, destination: %s", attribute, operator, split, value, destination), debug)
	return rule{
		attribute:   attribute,
		operator:    operator,
		value:       value,
		destination: destination,
	}
}

func parsePart(line string) part {
	debug := false
	split := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "{", ""), "}", ""), ",")
	log(fmt.Sprintf("  split: %v", split), debug)
	log(fmt.Sprintf("  x: %s", strings.Split(split[0], "=")[1]), debug)
	x := utils.Atoi(strings.Split(split[0], "=")[1])
	m := utils.Atoi(strings.Split(split[1], "=")[1])
	a := utils.Atoi(strings.Split(split[2], "=")[1])
	s := utils.Atoi(strings.Split(split[3], "=")[1])
	return part{x: x, m: m, a: a, s: s, pileName: "in"}
}

func log(s string, debug bool) {
	if debug {
		fmt.Println(s)
	}
}
