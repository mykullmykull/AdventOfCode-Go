package day

import (
	"goAoc2022/utils"
	"math"
)

type rock struct {
	points [][]int
}

func (r rock) getHeight() int {
	minRow := r.getMinRow(-1)
	maxRow := r.getMaxRow(-1)
	return maxRow - minRow + 1
}

func (r rock) getMinRow(col int) int {
	minRow := math.MaxInt
	for _, point := range r.points {
		if col < 0 || point[1] == col {
			minRow = utils.UpdateLeast(minRow, point[0])
		}
	}
	return minRow
}

func (r rock) getMaxRow(col int) int {
	maxRow := 0
	for _, point := range r.points {
		if col < 0 || point[1] == col {
			maxRow = utils.UpdateMost(maxRow, point[0])
		}
	}
	return maxRow
}

func (r rock) blowRock(wind string) rock {
	if r.cannotBlow(wind, columns{}) {
		return r
	}

	newPoints := make([][]int, len(r.points))
	copy(newPoints, r.points)
	for x := 0; x < len(r.points); x++ {
		newPoints[x] = r.movePoint(x, wind)
	}
	return rock{newPoints}
}

func (r rock) blowFallenRock(wind string, spaces columns) rock {
	if r.cannotBlow(wind, spaces) {
		return r
	}

	newPoints := make([][]int, len(r.points))
	copy(newPoints, r.points)
	for x := 0; x < len(r.points); x++ {
		newPoints[x] = r.movePoint(x, wind)
	}
	return rock{newPoints}
}

func (r rock) cannotBlow(wind string, spaces columns) bool {
	if wind != "<" && wind != ">" {
		panic("invalid wind" + wind)
	}

	if r.isAtTheEdge(wind) {
		return true
	}

	if r.isBlocked(wind, spaces) {
		return true
	}

	return false
}

func (r rock) isAtTheEdge(wind string) bool {
	for _, point := range r.points {
		col := point[1]
		if wind == "<" && col == 1 {
			return true
		}
		if wind == ">" && col == 7 {
			return true
		}
	}
	return false
}

func (r rock) isBlocked(wind string, towerCols columns) bool {
	if towerCols.isEmpty() {
		return false
	}

	for _, point := range r.points {
		row := point[0]
		col := point[1]
		if wind == "<" && towerCols.isBlocked(col-1, row) {
			return true
		}
		if wind == ">" && towerCols.isBlocked(col+1, row) {
			return true
		}
	}
	return false
}

func (r rock) movePoint(index int, wind string) []int {
	point := r.points[index]
	switch wind {
	case "<":
		if point[1] <= 1 {
			return point
		}
		return []int{point[0], point[1] - 1}
	case ">":
		if point[1] >= 7 {
			return point
		}
		return []int{point[0], point[1] + 1}
	default:
		panic("invalid wind" + wind)
	}
}

func (r rock) drop() rock {
	newPoints := make([][]int, len(r.points))
	for x, point := range r.points {
		newPoints[x] = []int{point[0] + 1, point[1]}
	}
	return rock{newPoints}
}
