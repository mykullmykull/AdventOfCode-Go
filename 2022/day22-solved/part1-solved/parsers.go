package day

func parse(grid []string, instructions string) board {
	b := board{
		grid:         grid,
		facingCol:    1,
		instructions: instructions,
	}
	walls := [][]int{}

	rowWidths := make([]int, len(grid))
	rowMins := make([]int, len(grid))
	rowMaxs := make([]int, len(grid))
	for r := range grid {
		rowMins[r] = -1
		rowMaxs[r] = -1
	}

	colHeights := make([]int, len(grid[0]))
	colMins := make([]int, len(grid[0]))
	colMaxs := make([]int, len(grid[0]))
	for c := range grid[0] {
		colMins[c] = -1
		colMaxs[c] = -1
	}

	for r, row := range grid {
		rowWalls := []int{}
		for c, cell := range row {
			if cell == '#' {
				rowWalls = append(rowWalls, c)
			}

			rowMin := rowMins[r]
			if rowMin == -1 && cell != ' ' {
				rowMins[r] = c
			}

			rowMax := rowMaxs[r]
			if rowMin > -1 && rowMax == -1 && cell == ' ' {
				rowMaxs[r] = c - 1
			}
			if c == len(row)-1 && rowMax == -1 {
				rowMaxs[r] = c
			}

			colMin := colMins[c]
			if colMin == -1 && cell != ' ' {
				colMins[c] = r
			}

			colMax := colMaxs[c]
			if colMin > -1 && colMax == -1 && cell == ' ' {
				colMaxs[c] = r - 1
			}
			if r == len(grid)-1 && colMax == -1 {
				colMaxs[c] = r
			}
		}
		rowWidths[r] = rowMaxs[r] - rowMins[r] + 1
		walls = append(walls, rowWalls)
	}
	for c := range grid[0] {
		colHeights[c] = colMaxs[c] - colMins[c] + 1
	}
	b.walls = walls
	b.rowWidths = rowWidths
	b.rowMins = rowMins
	b.rowMaxs = rowMaxs
	b.colHeights = colHeights
	b.colMins = colMins
	b.colMaxs = colMaxs
	b.col = b.rowMins[0]
	return b
}
