package day

import (
	"fmt"
	"goAoc2022/utils"
	"strings"
)

func printInput(input []string) {
	for _, line := range input {
		println(line)
	}
}

func printValves(names []string, valves map[string]valve) {
	for _, name := range names {
		v := valves[name]
		println(v.name, v.flow, utils.JoinArray(v.adjacentValves, ", "))
		printTargetsDist(v.targetsDist)
	}
}

func printTargetsDist(targetsDist map[string]int) {
	for k, v := range targetsDist {
		println("  ", k, v)
	}
}

func getSpacer(path []string) string {
	spacer := ""
	for x := 0; x < len(path); x++ {
		spacer += "  "
	}
	return spacer
}

func (s state) toString() string {
	return fmt.Sprintf("%d %s %b", s.timeLeft, s.position, s.openValves)
}

func (s state) toOpenStr() string {
	bin := fmt.Sprintf("%b", s.openValves)
	zerosToAdd := s.targetsCount - len(bin)
	return strings.Repeat("0", zerosToAdd) + bin
}

func (v valve) toString() string {
	str := fmt.Sprintf("%s %d %s", v.name, v.flow, utils.JoinArray(v.adjacentValves, ", "))
	for k, v := range v.targetsDist {
		str += fmt.Sprintf("\n  %s:%d", k, v)
	}
	return str
}

func printMap(nameToValve map[string]valve) {
	for _, v := range nameToValve {
		println(v.toString())
	}
}

func printPath(path []int, unopenedTargets []string) {
	targets := []string{}
	for _, index := range path {
		targets = utils.AppendToStrArray(targets, unopenedTargets[index])
	}
	println("path:", utils.JoinArray(targets, ", "))
}
