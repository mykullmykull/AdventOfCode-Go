package day

import (
	"fmt"
	"goAoc2024/utils"
)

type point struct {
	row int
	col int
}

type cheat struct {
	start point
	end   point
}

func main(baseMaze []string, target int, exact bool) int {
	baseRoute := mapBaseRoute(baseMaze)
	return countCheats(baseMaze, target, baseRoute, exact)
}

func mapBaseRoute(baseMaze []string) map[point]int {
	maze := make([]string, len(baseMaze))
	copy(maze, baseMaze)
	start := findPoint(maze, 'S')
	maze = updateGrid(maze, start, "O", true)
	baseRoute := make(map[point]int)
	baseRoute[start] = 0
	steps := 0
	r := 0
	c := 0
	for {
		if c >= len(maze[r]) {
			r++
			c = 0
		}
		if r >= len(maze) {
			r = 0
			c = 0
		}
		cell := maze[r][c]
		p := point{row: r, col: c}
		c++
		if cell == '#' || cell == 'O' {
			continue
		}
		if isNextStep(maze, p) {
			steps++
			baseRoute[p] = steps
			if cell == 'E' {
				return baseRoute
			}
			maze = updateGrid(maze, p, "O", true)
		}
	}
}

func isNextStep(maze []string, p point) bool {
	neighbors := getNeighbors(maze, p)
	for _, n := range neighbors {
		if maze[n.row][n.col] == 'O' {
			return true
		}
	}
	return false
}

func countCheats(maze []string, target int, baseRoute map[point]int, exact bool) int {
	cheats := map[cheat]int{}
	// consider starting a cheat at every point
	for r, row := range maze {
		for c, cell := range row {
			p := point{row: r, col: c}
			cheatStart := p
			if cell == '#' {
				err := error(nil)
				cheatStart, err = getSmallestNeighbor(maze, p, baseRoute)
				if err != nil {
					continue
				}
			}
			cheat := cheat{start: cheatStart}
			cheatStartStep := baseRoute[cheat.start]
			// consider ending the cheat at every point
			for r2, row2 := range maze {
				for c2, cell2 := range row2 {
					cheat.end = point{row: r2, col: c2}
					if cell2 == '#' {
						continue
					}
					cheatLength := utils.Abs(cheat.start.row-cheat.end.row) + utils.Abs(cheat.start.col-cheat.end.col)
					stepsAfterCheat := cheatStartStep + cheatLength
					savings := baseRoute[cheat.end] - stepsAfterCheat
					if cheatLength <= 20 && savings > 0 {
						cheats[cheat] = savings
					}
				}
			}
		}
	}
	cheatsMatchingTarget := 0
	for _, savings := range cheats {
		if exact && savings == target {
			cheatsMatchingTarget++
		}
		if !exact && savings >= target {
			cheatsMatchingTarget++
		}
	}
	return cheatsMatchingTarget
}

func getSmallestNeighbor(maze []string, p point, baseRoute map[point]int) (point, error) {
	neighbors := getNeighbors(maze, p)
	smallestStep := 1000000
	smallestNeighbor := point{-1, -1}
	hasFound := false
	for _, n := range neighbors {
		nSteps, ok := baseRoute[n]
		if !ok {
			continue
		}
		hasFound = true
		if nSteps < smallestStep {
			smallestStep = baseRoute[n]
			smallestNeighbor = n
		}
	}
	if !hasFound {
		return smallestNeighbor, fmt.Errorf("no neighboring steps found for %v", p)
	}
	return smallestNeighbor, nil
}
