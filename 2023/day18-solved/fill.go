package day18

type point struct {
	row int
	col int
}

func (p point) adjacentPoints(grid []string) []point {
	up := point{row: p.row - 1, col: p.col}
	dn := point{row: p.row + 1, col: p.col}
	lt := point{row: p.row, col: p.col - 1}
	rt := point{row: p.row, col: p.col + 1}
	return []point{up, dn, lt, rt}
}

func (p point) isValid(grid []string) bool {
	return p.row >= 0 &&
		p.row < len(grid) &&
		p.col >= 0 &&
		p.col < len(grid[p.row]) &&
		grid[p.row][p.col] != '#'
}

func (p point) fill(grid []string) []string {
	grid[p.row] = grid[p.row][:p.col] + "#" + grid[p.row][p.col+1:]
	return grid
}
