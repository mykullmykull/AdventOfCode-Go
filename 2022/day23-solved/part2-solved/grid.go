package day

import (
	"fmt"
	"strings"
)

func expand(grid []string) []string {
	emptyRow := strings.Repeat(".", len(grid[0]))
	if !isEmpty(grid, 0, -1) {
		grid = append([]string{emptyRow}, grid...)
	}
	if !isEmpty(grid, len(grid)-1, -1) {
		grid = append(grid, emptyRow)
	}
	if !isEmpty(grid, -1, 0) {
		for i := range grid {
			grid[i] = "." + grid[i]
		}
	}
	if !isEmpty(grid, -1, len(grid[0])-1) {
		for i := range grid {
			grid[i] = grid[i] + "."
		}
	}
	return grid
}

func isEmpty(grid []string, row int, col int) bool {
	if row == -1 {
		for _, line := range grid {
			if line[col] == '#' {
				return false
			}
		}
		return true
	}
	if col == -1 {
		for _, c := range grid[row] {
			if c == '#' {
				return false
			}
		}
		return true
	}
	panic("invalid row and col" + fmt.Sprint(row, col))
}

func countEmpty(grid []string) int {
	count := 0
	for _, row := range grid {
		for _, c := range row {
			if c == '.' || c == 'O' || c == 'X' {
				count++
			}
		}
	}
	return count
}

func lookAround(grid []string, y int, x int) []string {
	block := make([]string, 3)
	for z := -1; z < 2; z++ {
		block[z+1] = grid[y+z][x-1 : x+2]
	}
	return block
}

func shrink(grid []string) []string {
	removed := 0
	// Remove top row if empty
	if isEmpty(grid, 0, -1) {
		removed++
		grid = grid[1:]
	}
	// Remove bottom row if empty
	if isEmpty(grid, len(grid)-1, -1) {
		removed++
		grid = grid[:len(grid)-1]
	}
	// Remove first column if empty
	if isEmpty(grid, -1, 0) {
		removed++
		for i := range grid {
			grid[i] = grid[i][1:]
		}
	}
	// Remove last column if empty
	if isEmpty(grid, -1, len(grid[0])-1) {
		removed++
		for i := range grid {
			grid[i] = grid[i][:len(grid[i])-1]
		}
	}
	// Repeat until no more rows or columns can be removed
	if removed == 0 {
		return grid
	}

	return shrink(grid)
}
