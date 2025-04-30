package day

import "fmt"

func markEdges(grid []string) []string {
	for r, row := range grid {
		for c, _ := range row {
			if r == 0 || r == len(grid)-1 || c == 0 || c == len(row)-1 {
				grid = updateGrid(grid, point{r, c}, "*", false)
			}
		}
	}
	return grid
}

func findPoint(grid []string, val rune) point {
	for x, row := range grid {
		for y, cell := range row {
			if cell == val {
				return point{x, y}
			}
		}
	}
	panic("can't find point with value" + string(val))
}

func getNeighbors(grid []string, p point) []point {
	up := point{p.row - 1, p.col}
	dn := point{p.row + 1, p.col}
	lt := point{p.row, p.col - 1}
	rt := point{p.row, p.col + 1}
	return []point{up, dn, lt, rt}
}

func updateGrid(grid []string, w point, c string, force bool) []string {
	val := grid[w.row][w.col]
	if val == '*' {
		return grid
	}
	if !force && (val == 'S' || val == 'E') {
		return grid
	}
	grid[w.row] = grid[w.row][:w.col] + c + grid[w.row][w.col+1:]
	return grid
}

func lookup(grid []string, w point) string {
	if w.row < 0 || w.row >= len(grid) {
		return ""
	}
	if w.col < 0 || w.col >= len(grid[w.row]) {
		return ""
	}
	return string(grid[w.row][w.col])
}

func printMaze(maze []string) {
	for _, row := range maze {
		fmt.Println(row)
	}
	fmt.Println()
}
