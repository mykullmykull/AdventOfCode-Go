package day21

import "strings"

func runA(grid []string, steps int) int {
	subFarmSize := len(grid)
	grid = expandGrid(grid, calculateRowsOfSubFarms(steps, len(grid)))
	start := getStartingLocation(grid)
	locations := []location{start}
	for i := 0; i < steps; i++ {
		println("step", i)
		locations = getNextLocations(locations, grid)
	}
	printGrid(locations, grid, subFarmSize)
	return len(locations)
}

func getStartingLocation(grid []string) location {
	for r, row := range grid {
		for c, char := range row {
			if char == 'S' {
				return location{r, c}
			}
		}
	}
	return location{}
}

func getNextLocations(locations []location, grid []string) []location {
	locationsMap := map[location]bool{}
	for _, l := range locations {
		adjacentLocations := l.getNextLocations(grid)
		for _, nl := range adjacentLocations {
			locationsMap[nl] = true
		}
	}
	nextLocations := []location{}
	for k, _ := range locationsMap {
		nextLocations = append(nextLocations, k)
	}
	return nextLocations
}

func calculateRowsOfSubFarms(steps int, subFarmSize int) int {
	toEdge := subFarmSize / 2
	steps -= toEdge
	rowsOfSubFarms := 1
	for steps > 0 {
		rowsOfSubFarms += 2
		steps -= subFarmSize
	}
	return rowsOfSubFarms
}

func expandGrid(grid []string, rowsOfSubFarms int) []string {
	newGrid := make([]string, len(grid)*rowsOfSubFarms)
	sRow := getStartingLocation(grid).row
	sRow += len(grid) * (rowsOfSubFarms / 2)
	for i := 0; i < rowsOfSubFarms; i++ {
		for j, row := range grid {
			currentRow := i*len(grid) + j
			rowWithoutS := strings.Replace(row, "S", ".", -1)
			newRow := ""
			for k := 0; k < rowsOfSubFarms/2; k++ {
				newRow += rowWithoutS
			}
			if currentRow == sRow {
				newRow += row
			} else {
				newRow += rowWithoutS
			}
			for k := 0; k < rowsOfSubFarms/2; k++ {
				newRow += rowWithoutS
			}
			newGrid[currentRow] = newRow
		}
	}
	return newGrid
}

func printGrid(locations []location, grid []string, subFarmSize int) {
	for _, l := range locations {
		grid[l.row] = grid[l.row][:l.col] + "O" + grid[l.row][l.col+1:]
	}

	for i, row := range grid {
		newRow := ""
		counter := 0
		for _, char := range row {
			newRow += string(char)
			counter++
			if counter == subFarmSize {
				newRow += " "
				counter = 0
			}
		}
		grid[i] = newRow
	}

	println()
	rowIndex := 0
	for rowIndex < len(grid) {
		if rowIndex%subFarmSize == 0 {
			println()
		}
		println(grid[rowIndex])
		rowIndex++
	}
}
