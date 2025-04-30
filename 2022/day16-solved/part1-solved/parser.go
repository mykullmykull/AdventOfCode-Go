package day

import (
	"goAoc2022/utils"
	"strings"
)

func parse(input []string) (map[string]valve, []string) {
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
		flow := utils.StrToInt(split[1])
		adjacent := strings.Split(split[2], ", ")
		valves[name] = valve{name, flow, adjacent}
		names = utils.AppendToStrArray(names, name)
	}
	return valves, names
}
