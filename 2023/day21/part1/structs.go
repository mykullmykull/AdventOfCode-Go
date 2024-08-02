package day21

import "fmt"

type location struct {
	row int
	col int
}

func (l location) getNextLocations(grid []string) []location {
	locations := []location{}
	n := location{l.row - 1, l.col}
	s := location{l.row + 1, l.col}
	e := location{l.row, l.col + 1}
	w := location{l.row, l.col - 1}
	for _, loc := range []location{n, s, e, w} {
		if loc.isValid(grid) {
			locations = append(locations, loc)
		}

	}
	return locations
}

func (l location) isValid(grid []string) bool {
	return l.row >= 0 &&
		l.row < len(grid) &&
		l.col >= 0 &&
		l.col < len(grid[0]) &&
		grid[l.row][l.col] != '#'
}

func (l location) getNextLocations2(grid []string) []location {
	locations := []location{}
	n := location{l.row - 1, l.col}
	s := location{l.row + 1, l.col}
	e := location{l.row, l.col + 1}
	w := location{l.row, l.col - 1}
	for _, loc := range []location{n, s, e, w} {
		if loc.isValid2(grid) {
			locations = append(locations, loc)
		}

	}
	return locations
}

func (l location) isValid2(grid []string) bool {
	base := l.base(grid)
	return grid[base.row][base.col] != '#'
}

func (l location) toString() string {
	return fmt.Sprintf("row: %d, col: %d", l.row, l.col)
}

func (l location) address() int {
	return l.row*1000000000 + l.col
}

func (l location) base(grid []string) location {
	row := l.row % len(grid)
	if row < 0 {
		row += len(grid)
	}

	col := l.col % len(grid[0])
	if col < 0 {
		col += len(grid[0])
	}

	return location{row, col}
}
