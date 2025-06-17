package day

import (
	"goAoc2021/utils"
	"strings"
)

type scanner struct {
	x int
	y int
	z int
}

func main(input []string) int {
	scanners := parseScanners(input)
	maxDistance := 0
	for x := range scanners {
		s1 := scanners[x]
		for y := x + 1; y < len(scanners); y++ {
			s2 := scanners[y]
			distance := s1.manhattanDistance(s2)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}
	return maxDistance
}

func parseScanners(input []string) []scanner {
	scanners := []scanner{}
	for _, line := range input {
		println("~~ line", line)
		s := scanner{}
		split := strings.Split(line, ", ")
		println("~~ split", split)
		s.x = utils.Atoi(split[0])
		s.y = utils.Atoi(split[1])
		s.z = utils.Atoi(split[2])
		scanners = append(scanners, s)
	}
	return scanners
}

func (s scanner) manhattanDistance(other scanner) int {
	return utils.Abs(s.x-other.x) + utils.Abs(s.y-other.y) + utils.Abs(s.z-other.z)
}
