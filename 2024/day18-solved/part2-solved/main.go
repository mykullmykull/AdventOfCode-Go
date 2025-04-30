package day

import (
	"goAoc2024/utils"
	"strings"
)

func main(input []string, bytes int, size int) string {
	baseGrid := makeGrid(size)
	baseGrid = markCorruptedSpots(baseGrid, input, bytes)
	baseGrid[0] = markRow(baseGrid[0], 0, "0")
	for x, b := range input[bytes+1:] {
		println(bytes+x, b)
		baseGrid = markNewCorruption(baseGrid, b)
		grid := makeGrid(size)
		copy(grid, baseGrid)
		isBlocked := true
		for {
			// println()
			// printGrid(grid)
			reachedEnd := false
			grid, reachedEnd, isBlocked = takeSteps(grid)
			if reachedEnd || isBlocked {
				break
			}
		}
		if isBlocked {
			return b
		}
	}
	panic("No solution found")
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

func markNewCorruption(grid []string, point string) []string {
	split := strings.Split(point, ",")
	x := utils.Atoi(split[0])
	y := utils.Atoi(split[1])
	grid[y] = markRow(grid[y], x, "#")
	return grid
}

func takeSteps(grid []string) ([]string, bool, bool) {
	hasNewMark := false
	for row, line := range grid {
		for col := range line {
			if !isOnEdge(grid, row, col) {
				continue
			}
			if isFinished(grid, row, col) {
				return grid, true, false
			}
			grid[row] = markRow(grid[row], col, "1")
		}
	}
	for row, line := range grid {
		for col, spot := range line {
			if spot == '1' {
				hasNewMark = true
				grid[row] = markRow(grid[row], col, "0")
			}
		}
	}
	return grid, false, !hasNewMark
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
