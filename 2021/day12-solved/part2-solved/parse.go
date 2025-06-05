package day

import "strings"

func (c *caves) parseConnections(input []string) {
	connectionsMap := make(map[string]map[string]bool)
	for _, line := range input {
		caves := strings.Split(line, "-")

		// add cave2 to cave1's connections
		nextCaves := connectionsMap[caves[0]]
		if nextCaves == nil {
			nextCaves = make(map[string]bool)
		}
		nextCaves[caves[1]] = true
		connectionsMap[caves[0]] = nextCaves

		// add cave1 to cave2's connections
		nextCaves = connectionsMap[caves[1]]
		if nextCaves == nil {
			nextCaves = make(map[string]bool)
		}
		nextCaves[caves[0]] = true
		connectionsMap[caves[1]] = nextCaves
	}
	connections := make(map[string][]string)
	for cave, nextCaves := range connectionsMap {
		nextCavesList := make([]string, 0, len(nextCaves))
		for nextCave := range nextCaves {
			nextCavesList = append(nextCavesList, nextCave)
		}
		connections[cave] = nextCavesList
	}

	c.connections = connections
}
