package day

import (
	"goAoc2022/utils"
	"strings"
)

func parseTargets(input []string) (map[string]valve, []string, []string) {
	valves := map[string]valve{}
	names := []string{}
	for _, line := range input {
		trimmed := strings.Replace(
			strings.Replace(
				strings.Replace(
					strings.Replace(
						line,
						"Valve ", "", -1,
					),
					" has flow rate=", "|", -1,
				),
				"; tunnels lead to valves ", "|", -1,
			),
			"; tunnel leads to valve ", "|", -1,
		)
		split := strings.Split(trimmed, "|")
		name := split[0]
		names = utils.AppendToStrArray(names, name)
		flow := utils.StrToInt(split[1])
		adjacent := strings.Split(split[2], ", ")
		emptyTargetsDist := map[string]int{}
		valves[name] = valve{name, flow, adjacent, emptyTargetsDist}
	}
	targets := filterTargets(names, valves)
	valves = getDistances(valves, targets)
	return valves, names, targets
}

func filterTargets(names []string, valves map[string]valve) []string {
	filtered := []string{}
	for _, name := range names {
		if v, ok := valves[name]; ok {
			if v.flow > 0 {
				filtered = utils.AppendToStrArray(filtered, name)
			}
		}
	}
	return filtered
}

func getDistances(valves map[string]valve, targets []string) map[string]valve {
	for _, v := range valves {
		if !(v.name == "AA" || utils.ContainsStr(targets, v.name)) {
			continue
		}
		for _, t := range targets {
			v.targetsDist[t] = v.getDistanceTo([]string{v.name}, t, valves)
		}
	}
	return valves
}
