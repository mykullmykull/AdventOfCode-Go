package day24

import (
	"goAoc2023/utils"
	"strings"
)

// part 1
func parseStns(input []string) []stn {
	stones := []stn{}
	for _, line := range input {
		stones = append(stones, parseStn(line))
	}
	return stones
}

func parseStn(line string) stn {
	split := strings.Split(line, " @ ")
	positionsSplit := strings.Split(split[0], ", ")
	velocitiesSplit := strings.Split(split[1], ", ")
	return stn{
		start: point{
			x: float64(utils.Atoi(strings.Trim(positionsSplit[0], " "))),
			y: float64(utils.Atoi(strings.Trim(positionsSplit[1], " "))),
			z: float64(utils.Atoi(strings.Trim(positionsSplit[2], " "))),
		},
		dx: float64(utils.Atoi(strings.Trim(velocitiesSplit[0], " "))),
		dy: float64(utils.Atoi(strings.Trim(velocitiesSplit[1], " "))),
		dz: float64(utils.Atoi(strings.Trim(velocitiesSplit[2], " "))),
	}
}
