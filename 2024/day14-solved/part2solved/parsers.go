package day

import (
	"goAoc2024/utils"
	"strings"
)

func parse(input []string) []robot {
	robots := []robot{}
	for _, line := range input {
		robots = append(robots, parseRobot(line))
	}
	return robots
}

func parseRobot(line string) robot {
	line = strings.ReplaceAll(strings.ReplaceAll(line,
		"p=", ""),
		"v=", "")
	split := strings.Split(line, " ")
	posSplit := strings.Split(split[0], ",")
	pos := point{utils.Atoi(posSplit[0]), utils.Atoi(posSplit[1])}
	velSplit := strings.Split(split[1], ",")
	vel := point{utils.Atoi(velSplit[0]), utils.Atoi(velSplit[1])}
	return robot{pos, vel}
}
