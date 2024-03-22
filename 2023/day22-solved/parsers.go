package day22

import (
	"goAoc2023/utils"
	"strings"
)

func parseInput(input []string) []brick {
	bricks := []brick{}
	for _, line := range input {
		bricks = append(bricks, parseBrick(line))
	}
	return bricks
}

func parseBrick(line string) brick {
	startEndSplit := strings.Split(line, "~")
	startSplit := strings.Split(startEndSplit[0], ",")
	endSplit := strings.Split(startEndSplit[1], ",")
	xSpan := span{
		min: utils.Atoi(startSplit[0]),
		max: utils.Atoi(endSplit[0]),
	}
	ySpan := span{
		min: utils.Atoi(startSplit[1]),
		max: utils.Atoi(endSplit[1]),
	}
	zSpan := span{
		min: utils.Atoi(startSplit[2]),
		max: utils.Atoi(endSplit[2]),
	}

	return brick{x: xSpan, y: ySpan, z: zSpan}
}
