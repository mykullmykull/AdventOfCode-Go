package day

import (
	"goAoc2022/utils"
	"math"
)

func main(input []string) int {
	// parse valves
	valves, _ := parse(input)
	for _, v := range valves {
		println(v.toString())
	}

	// initiate paths
	paths := []path{{
		location: valves["AA"],
		time:     0,
		released: 0,
		history:  map[string]bool{},
	}}

	// get targets
	targets := getTargets(valves)

	max := 0
	for {
		newPaths := []path{}
		for _, p := range paths {
			max = utils.UpdateMost(max, p.released)
			for _, t := range targets {
				toAdd, err := p.moveToTarget(t, valves)
				if err == nil {
					newPaths = appendToPaths(newPaths, toAdd)
				}
			}
		}
		if len(newPaths) == 0 {
			break
		}
		paths = filter(newPaths, targets)
	}
	return max
}

func getTargets(valves map[string]valve) []string {
	filtered := []string{}
	for k, v := range valves {
		if v.flow == 0 {
			continue
		}
		filtered = utils.AppendToStrArray(filtered, k)
	}
	return filtered
}

func printPaths(paths []path) {
	for _, p := range paths {
		println(p.toString())
	}
}

func updateMaxHistory(history map[string]bool) []string {
	newHistory := []string{}
	for k := range history {
		newHistory = append(newHistory, k)
	}
	return newHistory
}

func filter(paths []path, targets []string) []path {
	filtered := []path{}
	for _, t := range targets {
		maxReleased := 0
		minTime := math.MaxInt
		for _, p := range paths {
			if t == p.location.name {
				if p.released > maxReleased && p.time <= minTime {
					maxReleased = p.released
					minTime = p.time
				}
			}
		}
		for _, p := range paths {
			if t == p.location.name {
				if p.released >= maxReleased || p.time < minTime {
					filtered = append(filtered, p)
					continue
				}
			}
		}
	}
	return filtered
}
