package day

import (
	"strconv"
	"strings"
)

func main(input []string, iterations int) int {
	enhancementAlgorithm := input[0]
	grid := input[2:]

	for x := 0; x < iterations; x++ {
		grid = enhance(grid, enhancementAlgorithm, x)
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

func enhance(grid []string, algorithm string, iteration int) []string {
	char := getBorderChar(algorithm, iteration)
	grid = addBorder(grid, char)
	grid = addBorder(grid, char)

	// Build new grid of the same size
	newGrid := make([]string, len(grid))

	// Loop from row 1 to len(grid)-2 and col 1 to len(grid[row])-2
	for row := 1; row < len(grid)-1; row++ {
		var newRow strings.Builder
		for col := 1; col < len(grid[row])-1; col++ {
			// Grab 9 characters in a 3x3 square centered on current cell
			square := get3x3Square(grid, row, col)

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

func addBorder(grid []string, char string) []string {
	blankRow := strings.Repeat(char, len(grid[0])+2)
	grid = append([]string{blankRow}, grid...)

	for row := range grid {
		if row == 0 {
			continue
		}
		gridRow := grid[row]
		gridRow = char + gridRow + char
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

func getBorderChar(algorithm string, iteration int) string {
	if iteration == 0 {
		return "."
	}

	oddChar := algorithm[0]
	evenChar := oddChar
	if oddChar == '#' {
		fullyLitX, _ := strconv.ParseInt("111111111", 2, 64)
		evenChar = algorithm[fullyLitX]
	}
	var char byte
	if iteration%2 == 0 {
		char = evenChar
	}
	if iteration%2 == 1 {
		char = oddChar
	}
	return string(char)
}
