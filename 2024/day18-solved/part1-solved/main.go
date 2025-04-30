package day

import (
	"goAoc2024/utils"
	"strings"
)

func main(input []string, bytes int, size int) int {
	grid := makeGrid(size)
	grid = markCorruptedSpots(grid, input, bytes)
	step := 0
	grid[0] = markRow(grid[0], 0, "0")
	for {
		println()
		println(step)
		printGrid(grid)
		step++
		isFinished := false
		grid, isFinished = takeSteps(grid)
		if isFinished {
			return step
		}
	}
}

func makeGrid(size int) []string {
	grid := make([]string, size)
	row := strings.Repeat(".", size)
	for i := 0; i < size; i++ {
		grid[i] = row
	}
	return grid
}

func markCorruptedSpots(grid []string, input []string, bytes int) []string {
	for i, point := range input {
		if i >= bytes {
			break
		}
		split := strings.Split(point, ",")
		x := utils.Atoi(split[0])
		y := utils.Atoi(split[1])
		grid[y] = markRow(grid[y], x, "#")
	}
	return grid
}

func markRow(row string, x int, mark string) string {
	after := ""
	if x < len(row)-1 {
		after = row[x+1:]
	}
	return row[:x] + mark + after
}

func takeSteps(grid []string) ([]string, bool) {
	for row, line := range grid {
		for col := range line {
			if !isOnEdge(grid, row, col) {
				continue
			}
			if isFinished(grid, row, col) {
				return grid, true
			}
			grid[row] = markRow(grid[row], col, "1")
		}
	}
	for row, line := range grid {
		for col, spot := range line {
			if spot == '1' {
				grid[row] = markRow(grid[row], col, "0")
			}
		}
	}
	return grid, false
}

func isOnEdge(grid []string, row int, col int) bool {
	if grid[row][col] != '.' {
		return false
	}
	if row > 0 && grid[row-1][col] == '0' {
		return true
	}
	if col < len(grid)-1 && grid[row][col+1] == '0' {
		return true
	}
	if row < len(grid)-1 && grid[row+1][col] == '0' {
		return true
	}
	if col > 0 && grid[row][col-1] == '0' {
		return true
	}
	return false
}

func isFinished(grid []string, row int, col int) bool {
	return row == len(grid)-1 && col == len(grid)-1
}

func printGrid(grid []string) {
	for _, row := range grid {
		println(row)
	}
}
