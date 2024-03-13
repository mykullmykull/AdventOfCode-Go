package day18

import (
	"fmt"
	"goAoc2023/utils"
	"strings"
)

func runA(input []string) int {
	debug := false
	edges := [][]int{}
	row := 0
	col := 0
	for _, line := range input {
		split := strings.Split(line, " ")
		direction := split[0]
		distance := utils.Atoi(split[1])
		dRow, dCol := parseDirection(direction)
		e := edge{
			startRow: row,
			startCol: col,
			dRow:     dRow,
			dCol:     dCol,
			distance: distance,
		}
		log(fmt.Sprintf("edge: %+v", e), debug)
		edges, row, col = e.markPoints(edges)
	}

	grid := drawGrid(edges)
	printGrid(grid)
	count := countEdges(edges)
	count = fillAndCount(grid, count)
	return count
}

func parseDirection(direction string) (int, int) {
	switch direction {
	case "U":
		return -1, 0
	case "D":
		return 1, 0
	case "L":
		return 0, -1
	case "R":
		return 0, 1
	}
	panic("invalid direction")
}

func drawGrid(edges [][]int) []string {
	grid := []string{}
	for _, eRow := range edges {
		gRow := ""
		maxGCol := eRow[len(eRow)-1] + 1
		eRowCursor := 0
		nextEdgeCol := eRow[eRowCursor]
		for gCol := 0; gCol < maxGCol; gCol += 1 {
			if gCol == nextEdgeCol {
				gRow += "#"
				if eRowCursor < len(eRow)-1 {
					eRowCursor += 1
					nextEdgeCol = eRow[eRowCursor]
				}
				continue
			}
			gRow += "."
		}
		grid = append(grid, gRow)
	}
	return grid
}

func countEdges(edges [][]int) int {
	count := 0
	for _, eRow := range edges {
		count += len(eRow)
	}
	return count
}

func fillAndCount(grid []string, count int) int {
	debug := false
	startingPoint := point{row: 1, col: 0}
	for col := 0; col < len(grid[0]); col += 1 {
		if grid[0][col] == '#' {
			startingPoint.col = col + 1
			break
		}
	}

	frontier := []point{startingPoint}
	for {
		// log(fmt.Sprintf("count: %d", count), debug)
		if len(frontier) == 0 {
			break
		}
		currentPoint := frontier[0]
		log(fmt.Sprintf("  currentPoint: %v", currentPoint), debug)
		frontier = frontier[1:]
		if !currentPoint.isValid(grid) {
			// log("	invalid", debug)
			continue
		}
		grid = currentPoint.fill(grid)
		count += 1
		frontier = append(frontier, currentPoint.adjacentPoints(grid)...)
	}

	return count
}

// ---------------------------------------- debugging utils ----------------------------------------

func log(s string, debug bool) {
	if debug {
		fmt.Println(s)
	}
}

func printEdgesPoints(edges [][]int) {
	for _, row := range edges {
		for _, col := range row {
			print(col, " ")
		}
		println()
	}
	println()
}

func printGrid(grid []string) {
	println()
	for _, row := range grid {
		println(row)
	}
	println()
}
