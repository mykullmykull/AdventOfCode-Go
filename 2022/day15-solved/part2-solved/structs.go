package day

import (
	"fmt"
	"goAoc2022/utils"
)

type point struct {
	col int
	row int
}

type sensor struct {
	p             point
	closestBeacon point
	distance      int
}

type rowCoverage struct {
	min int
	max int
}

func (p point) distanceFrom(p2 point) int {
	return utils.AbsForInt(p.row-p2.row) + utils.AbsForInt(p.col-p2.col)
}

func (s sensor) toString() string {
	return fmt.Sprintf("p %v closestBeacn %v distance %v", s.p, s.closestBeacon, s.distance)
}

func (s sensor) rowMinMax(row int, rowMax int) (int, int) {
	remainingDistance := s.distance - utils.AbsForInt(row-s.p.row)
	if remainingDistance < 0 {
		return 0, 0
	}
	min := s.p.col - remainingDistance
	max := s.p.col + remainingDistance
	if min < 0 {
		min = 0
	}
	if max > rowMax {
		max = rowMax
	}
	return min, max
}

func (c rowCoverage) combineIfOverlaps(c2 rowCoverage) []rowCoverage {
	overlaps := c.min <= c2.max && c.max >= c2.min
	if overlaps {
		if c2.min < c.min {
			c.min = c2.min
		}
		if c2.max > c.max {
			c.max = c2.max
		}
		return []rowCoverage{c}
	}
	return []rowCoverage{c, c2}
}
