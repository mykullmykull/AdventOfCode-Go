package day

import "goAoc2024/utils"

func run(grid []string, start point, end point) int {
	grid = updateGrid(grid, start, "0", true)
	time := 0
	for {
		// printGrid(grid)
		if isFinished(grid, end) {
			return time
		}
		grid = advance(grid)
		time++
	}
}

func isFinished(grid []string, end point) bool {
	for _, n := range getNeighbors(grid, end) {
		if grid[n.row][n.col] == '0' {
			return true
		}
	}
	return false
}

func advance(grid []string) []string {
	newGrid := make([]string, len(grid))
	copy(newGrid, grid)
	for r, row := range grid {
		for c, val := range row {
			if val == '#' || val == '*' {
				continue
			}
			p := point{r, c}
			for _, n := range getNeighbors(grid, p) {
				nVal := grid[n.row][n.col]
				if nVal == '0' {
					newGrid = updateGrid(newGrid, p, "0", true)
					break
				}
			}
		}
	}
	return newGrid
}

func runWithCheat(grid []string, start point, end point) int {
	cheatStart := findPoint(grid, '1')
	cheatEnd := findPoint(grid, '2')

	gridCopy := make([]string, len(grid))
	copy(gridCopy, grid)
	t1 := run(gridCopy, start, cheatStart)
	t2 := utils.Abs(cheatStart.row-cheatEnd.row) + utils.Abs(cheatStart.col-cheatEnd.col)
	copy(gridCopy, grid)
	t3 := run(gridCopy, cheatEnd, end)
	return t1 + t2 + t3 + 1
}
