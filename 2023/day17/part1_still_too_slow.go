package day17

import (
	"fmt"
	"goAoc2023/utils"
)

func (p *point) equals(p2 point) bool {
	return p.row == p2.row && p.col == p2.col
}

func runA_still_too_slow(grid []string) int {
	debug := true
	p := point{0, 0}
	heatLoss := 0
	path := ""
	isValid := true

	finish := point{len(grid) - 1, len(grid[0]) - 1}
	lowestHeatLossToFinish := -1
	bestPath := ""
	for {
		// log(fmt.Sprintf("path: %s, p: %v", path, p), debug && isValid)
		// log(fmt.Sprintf("path: %s, p: %v, heatLoss: %d, isValid: %v", path, p, heatLoss, isValid), debug && isValid)
		p, path = addNextMove(p, path, isValid)
		isValid = arePointAndPathValid(p, path, grid)
		if !isValid {
			if len(path) > 0 {
				continue
			}
			break
		}

		heatLoss += utils.Atoi(string(getRuneFromGrid(grid, p)))
		if p.equals(finish) && (lowestHeatLossToFinish == -1 || heatLoss < lowestHeatLossToFinish) {
			log(fmt.Sprintf("path: %s, p: %v, heatLoss: %d, lowestHeatLossToFinish: %d", path, p, heatLoss, lowestHeatLossToFinish), debug)
			lowestHeatLossToFinish = heatLoss
			bestPath = path
		}
	}
	log(fmt.Sprintf("bestPath: %s", bestPath), debug)
	return lowestHeatLossToFinish
}

func addNextMove(p point, path string, isValid bool) (point, string) {
	return p, path
}

func arePointAndPathValid(p point, path string, grid []string) bool {
	return p.row >= 0 && p.row < len(grid) &&
		p.col >= 0 && p.col < len(grid[0]) &&
		getRecentRepeatedMoves(path) <= 3 &&
		!movedBackward(path) &&
		!loops(p, path)
}

func movedBackward(path string) bool {
	if len(path) <= 1 {
		return false
	}

	m1 := path[len(path)-1]
	m2 := path[len(path)-2]
	return (m1 == '^' && m2 == 'v') ||
		(m1 == '>' && m2 == '<') ||
		(m1 == 'v' && m2 == '^') ||
		(m1 == '<' && m2 == '>')
}

func getRuneFromGrid(grid []string, p point) rune {
	return rune(grid[p.row][p.col])
}
