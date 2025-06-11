package day

import (
	"goAoc2021/utils"
	"strings"
)

func parse(input []string) (map[dot]bool, []fold) {
	dots := []dot{}
	folds := []fold{}
	finishedDots := false
	for _, line := range input {
		if line == "" {
			finishedDots = true
			continue
		}
		if !finishedDots {
			dots = append(dots, parseDot(line))
			continue
		}
		folds = append(folds, parseFold(line))
	}
	mapDots := make(map[dot]bool)
	for _, d := range dots {
		mapDots[d] = true
	}
	return mapDots, folds
}

func parseDot(line string) dot {
	split := strings.Split(line, ",")
	x := utils.Atoi(split[0])
	y := utils.Atoi(split[1])
	return dot{x, y}
}

func parseFold(line string) fold {
	str := strings.ReplaceAll(line, "fold along ", "")
	split := strings.Split(str, "=")
	dimension := split[0]
	distance := utils.Atoi(split[1])
	return fold{dimension, distance}
}
