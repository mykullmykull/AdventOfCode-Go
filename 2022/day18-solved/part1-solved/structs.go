package day

import (
	"goAoc2022/utils"
)

type cube struct {
	x int
	y int
	z int
}

func (c cube) countNeighbors(cubes []cube) int {
	count := 0
	for _, c2 := range cubes {
		if c.x == c2.x && c.y == c2.y && c.z == c2.z {
			continue
		}

		if (c.x == c2.x && c.y == c2.y && utils.AbsForInt(c.z-c2.z) == 1) ||
			(c.x == c2.x && c.z == c2.z && utils.AbsForInt(c.y-c2.y) == 1) ||
			(c.y == c2.y && c.z == c2.z && utils.AbsForInt(c.x-c2.x) == 1) {
			count++
		}
	}
	return count
}

func (c cube) countFreeSides(cubes []cube) int {
	return 6 - c.countNeighbors(cubes)
}
