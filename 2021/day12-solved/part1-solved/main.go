package day

import (
	"strings"
)

func main(input []string) int {
	connections := parseConnections(input)
	// fmt.Println(connections)
	finals := make(map[string]bool)
	partials := [][]string{{"start"}}
	for len(partials) > 0 {
		path := partials[0]
		partials = partials[1:]
		lastCave := path[len(path)-1]
		for cave := range connections[lastCave] {
			newPath := make([]string, len(path))
			copy(newPath, path)
			if hasRepeatedSmallCave(path, cave) {
				continue
			}
			newPath = append(newPath, cave)
			if cave == "end" {
				finals[strings.Join(path, ",")] = true
				continue
			}
			partials = append(partials, newPath)
		}
	}
	return len(finals)
}

func parseConnections(input []string) map[string]map[string]bool {
	connections := make(map[string]map[string]bool)
	for _, line := range input {
		caves := strings.Split(line, "-")

		// add cave2 to cave1's connections
		nextCaves := connections[caves[0]]
		if nextCaves == nil {
			nextCaves = make(map[string]bool)
		}
		nextCaves[caves[1]] = true
		connections[caves[0]] = nextCaves

		// add cave1 to cave2's connections
		nextCaves = connections[caves[1]]
		if nextCaves == nil {
			nextCaves = make(map[string]bool)
		}
		nextCaves[caves[0]] = true
		connections[caves[1]] = nextCaves
	}
	return connections
}

func hasRepeatedSmallCave(path []string, cave string) bool {
	if strings.ToLower(cave) != cave {
		return false
	}

	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}
