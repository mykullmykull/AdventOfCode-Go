package day

import (
	"fmt"
	"math"
)

type point struct {
	row int
	col int
}

type sensor struct {
	p             point
	closestBeacon point
	distance      int
}

func (p point) distanceFrom(p2 point) int {
	return int(math.Abs(float64(p.row)-float64(p2.row))) + int(math.Abs(float64(p.col)-float64(p2.col)))
}

func (s sensor) toString() string {
	return fmt.Sprintf("p %v closestBeacn %v distance %v", s.p, s.closestBeacon, s.distance)
}
