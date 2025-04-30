package day

func parse(grid []string, instructions string) board {
	b := board{
		grid:         grid,
		instructions: instructions,
		p:            point{row: 0, col: getFirstCol(grid[0]), side: tp},
		isTesting:    instructions == testInstructions,
	}
	b.p = b.p.changeDirection(rt)
	for row, line := range grid {
		for col, char := range line {
			if char != '#' {
				continue
			}
			wallP := point{row: row, col: col}
			wallP.side = b.getSide(wallP)
			b.walls = append(b.walls, wallP)
		}
	}
	return b
}

func getFirstCol(row string) int {
	for c, char := range row {
		if char == '.' {
			return c
		}
	}
	panic("couldn't find the first . in " + row)
}
