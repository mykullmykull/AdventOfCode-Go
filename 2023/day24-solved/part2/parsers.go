package day24Part2

import (
	"goAoc2023/utils"
	"strings"
)

// part 2
func parseStones(input []string) []stone {
	stones := []stone{}
	for _, line := range input {
		stones = append(stones, parseStone(line))
	}
	return stones
}

func parseStone(line string) stone {
	split := strings.Split(line, " @ ")
	positionsSplit := strings.Split(split[0], ", ")
	velocitiesSplit := strings.Split(split[1], ", ")
	return stone{
		x:  float64(utils.Atoi(strings.Trim(positionsSplit[0], " "))),
		y:  float64(utils.Atoi(strings.Trim(positionsSplit[1], " "))),
		z:  float64(utils.Atoi(strings.Trim(positionsSplit[2], " "))),
		dx: float64(utils.Atoi(strings.Trim(velocitiesSplit[0], " "))),
		dy: float64(utils.Atoi(strings.Trim(velocitiesSplit[1], " "))),
		dz: float64(utils.Atoi(strings.Trim(velocitiesSplit[2], " "))),
	}
}
