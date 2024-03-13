package day18

import (
	"fmt"
)

type edge struct {
	startRow int
	startCol int
	dRow     int
	dCol     int
	distance int
}

func (e edge) markPoints(grid [][]int) ([][]int, int, int) {
	row := e.startRow
	col := e.startCol
	for i := 0; i < e.distance; i++ {
		row += e.dRow
		col += e.dCol
		grid, row, col = expandGrid(grid, row, col)
		grid = markPoint(grid, row, col)
	}

	return grid, row, col
}

func expandGrid(grid [][]int, row, col int) ([][]int, int, int) {
	if len(grid) <= row {
		for i := len(grid); i <= row; i++ {
			grid = append(grid, []int{})
		}
	}

	if row < 0 {
		for i := 0; i < -row; i++ {
			row++
			grid = append([][]int{{}}, grid...)
		}
	}
	return grid, row, col
}

func markPoint(grid [][]int, row, col int) [][]int {
	if len(grid[row]) == 0 {
		grid[row] = append(grid[row], col)
		return grid
	}

	colAdded := false
	for i, c := range grid[row] {
		if c == col {
			return grid
		}

		if c > col {
			grid[row] = append(grid[row][:i], append([]int{col}, grid[row][i:]...)...)
			colAdded = true
			break
		}
	}

	if !colAdded {
		grid[row] = append(grid[row], col)
	}

	return grid
}

func (e edge) minCol() int {
	if e.dCol < 0 {
		return e.startCol + (e.dCol * e.distance)
	}
	return e.startCol
}

func (e edge) maxCol() int {
	if e.dCol > 0 {
		return e.startCol + (e.dCol * e.distance)
	}
	return e.startCol
}

func (e edge) minRow() int {
	if e.dRow < 0 {
		return e.startRow + (e.dRow * e.distance)
	}
	return e.startRow
}

func (e edge) maxRow() int {
	if e.dRow > 0 {
		return e.startRow + (e.dRow * e.distance)
	}
	return e.startRow
}

func (e edge) toSpan() span {
	return span{
		minRow: e.minRow(),
		maxRow: e.maxRow(),
		minCol: e.minCol(),
		maxCol: e.maxCol(),
	}
}

func (e edge) toString() string {
	return fmt.Sprintf("startRow: %d, startCol: %d, dRow: %d, dCol: %d, distance: %d", e.startRow, e.startCol, e.dRow, e.dCol, e.distance)
}
