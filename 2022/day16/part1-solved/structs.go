package day

import "fmt"

type valve struct {
	name           string
	flow           int
	adjacentValves []string
}

type path struct {
	location valve
	time     int
	released int
	history  map[string]bool
}

func (v valve) toString() string {
	return fmt.Sprintf("name: %s, flow: %d, adjacent: %v", v.name, v.flow, v.adjacentValves)
}

func (p path) toString() string {
	return fmt.Sprintf("time %d location %s released %d, history: %v", p.time, p.location.name, p.released, p.history)
}

func (v valve) stepsFrom(start string, valves map[string]valve) int {
	paths := []path{{
		location: valves[start],
		time:     0,
		history:  map[string]bool{},
	}}
	paths[0].history[start] = true
	for {
		newPaths := []path{}
		for _, p := range paths {
			if p.location.name == v.name {
				return p.time
			}
			pathsToAdd := p.move(valves)
			for _, newPath := range pathsToAdd {
				newPaths = appendToPaths(newPaths, newPath)
			}
		}
		paths = newPaths
	}
}

func appendToPaths(h []path, toAdd path) []path {
	copyOfArray := make([]path, len(h))
	copy(copyOfArray, h)
	copyOfArray = append(copyOfArray, toAdd)
	return copyOfArray
}

func (p path) move(valves map[string]valve) []path {
	newPaths := []path{}
	for _, adj := range p.location.adjacentValves {
		if p.history[adj] {
			continue
		}
		p.history[adj] = true
		newPath := path{
			location: valves[adj],
			time:     p.time + 1,
			history:  p.history,
		}
		newPaths = appendToPaths(newPaths, newPath)
	}
	return newPaths
}

func (p path) moveToTarget(target string, valves map[string]valve) (path, error) {
	if p.history[target] {
		return p, fmt.Errorf("path %s already visited %s", p.toString(), target)
	}
	v := valves[target]
	steps := v.stepsFrom(p.location.name, valves)
	if p.time+steps > 30 {
		return p, fmt.Errorf("time %d + steps %d would exceed 30", p.time, steps)
	}
	newTime := p.time + v.stepsFrom(p.location.name, valves) + 1
	newReleased := p.released + v.flow*(30-newTime)
	newHistory := map[string]bool{}
	for k := range p.history {
		newHistory[k] = true
	}
	newHistory[target] = true
	return path{
		time:     newTime,
		location: v,
		released: newReleased,
		history:  newHistory,
	}, nil
}

func (p path) matchesHistory(history []string) bool {
	isMissingTarget := false
	for _, t := range history {
		if !p.history[t] {
			isMissingTarget = true
			break
		}
	}
	return len(p.history) == len(history) && !isMissingTarget
}
