package day

import (
	"fmt"
)

type point struct {
	frq string
	row int
	col int
}

type pair struct {
	a point
	b point
}

func main(grid []string) int {
	count := 0
	antennas := getAntennas(grid)
	for x, a := range antennas {
		for _, b := range antennas[x+1:] {
			if a.frq != b.frq {
				continue
			}

			println()
			p := pair{a, b}
			antinodes := p.getAntinodes(grid)
			for _, antinode := range antinodes {
				if antinode.isUnique(grid) {
					grid = antinode.updateGrid(grid)
					count++
				}
			}
			printGrid(grid)
			println("count:", count)
		}
	}

	return count
}

func getAntennas(grid []string) []point {
	antennas := []point{}
	for r, row := range grid {
		for c, cell := range row {
			if cell == '.' {
				continue
			}
			antennas = append(antennas, point{frq: fmt.Sprintf("%c", cell), row: r, col: c})
		}
	}
	return antennas
}

func (p pair) getAntinodes(grid []string) []point {
	antinodes := []point{}
	rowGap := p.getRowGap()
	colGap := p.getColGap()
	addToA := point{frq: p.a.frq, row: p.a.row + rowGap, col: p.a.col + colGap}
	addToB := point{frq: p.a.frq, row: p.b.row + rowGap, col: p.b.col + colGap}
	subFromA := point{frq: p.a.frq, row: p.a.row - rowGap, col: p.a.col - colGap}
	subFromB := point{frq: p.a.frq, row: p.b.row - rowGap, col: p.b.col - colGap}

	for _, point := range []point{addToA, addToB, subFromA, subFromB} {
		if point == p.a || point == p.b {
			continue
		}
		if point.isInGrid(grid) {
			antinodes = append(antinodes, point)
		}
	}
	return antinodes
}

func (p pair) getRowGap() int {
	return p.a.row - p.b.row
}

func (p pair) getColGap() int {
	return p.a.col - p.b.col
}

func (p point) isInGrid(grid []string) bool {
	return p.row >= 0 &&
		p.row < len(grid) &&
		p.col >= 0 &&
		p.col < len(grid[0])
}

func (p point) updateGrid(grid []string) []string {
	newRow := grid[p.row][:p.col] + "#" + grid[p.row][p.col+1:]
	grid[p.row] = newRow
	return grid
}

func (p point) isUnique(grid []string) bool {
	return grid[p.row][p.col] != '#'
}
