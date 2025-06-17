package day

import (
	"strconv"
	"strings"
)

func main(input []string) int {
	enhancementAlgorithm := input[0]
	grid := input[2:]

	for range 2 {
		grid = enhanceGrid(grid, enhancementAlgorithm)
	}

	// Count lit pixels (#)
	count := 0
	for _, row := range grid {
		for _, char := range row {
			if char == '#' {
				count++
			}
		}
	}

	return count
}

func enhanceGrid(grid []string, algorithm string) []string {
	// Add border of '..' around the grid
	gridWithBorder := addBorder(addBorder(grid))

	// Build new grid of the same size
	newGrid := make([]string, len(gridWithBorder))

	// Loop from row 1 to len(grid)-2 and col 1 to len(grid[row])-2
	for row := 1; row < len(gridWithBorder)-1; row++ {
		var newRow strings.Builder
		for col := 1; col < len(gridWithBorder[row])-1; col++ {
			// Grab 9 characters in a 3x3 square centered on current cell
			square := get3x3Square(gridWithBorder, row, col)

			// Convert to binary (. = 0, # = 1)
			binaryStr := ""
			for _, char := range square {
				if char == '#' {
					binaryStr += "1"
				} else {
					binaryStr += "0"
				}
			}

			// Convert to decimal index
			index, _ := strconv.ParseInt(binaryStr, 2, 64)

			// Apply char at decimal index in enhancement algorithm
			newChar := algorithm[index]
			newRow.WriteByte(newChar)
		}
		newGrid[row] = newRow.String()
	}

	// Remove the border rows (first and last) that we didn't process
	result := make([]string, len(newGrid)-2)
	for i := 1; i < len(newGrid)-1; i++ {
		result[i-1] = newGrid[i]
	}

	return result
}

func addBorder(grid []string) []string {
	blankRow := strings.Repeat(".", len(grid[0])+2)
	grid = append([]string{blankRow}, grid...)

	for row := range grid {
		if row == 0 {
			continue
		}
		gridRow := grid[row]
		gridRow = "." + gridRow + "."
		grid[row] = gridRow
	}

	grid = append(grid, blankRow)
	return grid
}

func get3x3Square(grid []string, centerRow, centerCol int) string {
	square := ""
	square += grid[centerRow-1][centerCol-1 : centerCol+2] // Top row
	square += grid[centerRow][centerCol-1 : centerCol+2]   // Middle row
	square += grid[centerRow+1][centerCol-1 : centerCol+2] // Bottom row

	return square
}
