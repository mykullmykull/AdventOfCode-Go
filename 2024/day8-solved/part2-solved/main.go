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

			p := pair{a, b}
			antinodes := p.getAntinodes(grid)
			for _, antinode := range antinodes {
				if antinode.isUnique(grid) {
					grid = antinode.updateGrid(grid)
					count++
				}
			}
			println()
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
	antinodes := []point{p.a}
	rowGap := p.getRowGap()
	colGap := p.getColGap()

	// add to a
	r := p.a.row
	c := p.a.col
	for {
		r += rowGap
		c += colGap
		candidate := point{frq: p.a.frq, row: r, col: c}
		if candidate.isInGrid(grid) {
			antinodes = append(antinodes, candidate)
			continue
		}
		break
	}

	// subtract from a
	r = p.a.row
	c = p.a.col
	for {
		r -= rowGap
		c -= colGap
		candidate := point{frq: p.a.frq, row: r, col: c}
		if candidate.isInGrid(grid) {
			antinodes = append(antinodes, candidate)
			continue
		}
		break
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
