package day21

import "fmt"

type grid struct {
	grid []string
}

type location struct {
	r int
	c int
}

func (g grid) getStart() location {
	for r, row := range g.grid {
		for c, cell := range row {
			if cell == 'S' {
				return location{r, c}
			}
		}
	}
	return location{-1, -1}
}

func (g grid) getSpot(l location) string {
	reducedL := l.reduce(g)
	value := g.grid[reducedL.r][reducedL.c : reducedL.c+1]
	return value
}

func (g grid) updateSpot(l location) grid {
	row := g.grid[l.r]
	newRow := row[:l.c] + "O" + row[l.c+1:]
	g.grid[l.r] = newRow
	return g
}

func (g grid) printGrid(locations []location) {
	newGrid := make([]string, len(g.grid))
	copy(newGrid, g.grid)
	copyGrid := grid{newGrid}
	for _, loc := range locations {
		copyGrid = copyGrid.updateSpot(loc)
	}

	for _, row := range copyGrid.grid {
		fmt.Println(row)
	}
}

func (l location) reduce(g grid) location {
	horizontalGridLength := len(g.grid[0])
	verticalGridLength := len(g.grid)
	for {
		if l.r >= 0 {
			break
		}
		l.r = verticalGridLength + l.r
	}

	for {
		if l.c >= 0 {
			break
		}
		l.c = horizontalGridLength + l.c
	}

	r := l.r % verticalGridLength
	c := l.c % horizontalGridLength
	return location{r, c}
}

func (l location) equals(l2 location) bool {
	return l.r == l2.r && l.c == l2.c
}

func (l location) isValid(g grid, prev map[int]map[int]int) bool {
	if prev[l.r] != nil && prev[l.r][l.c] > 0 {
		return false
	}

	return g.getSpot(l) != "#"
}

func (l location) moveOne(g grid, prev map[int]map[int]int) []location {
	n := location{l.r - 1, l.c}
	s := location{l.r + 1, l.c}
	e := location{l.r, l.c + 1}
	w := location{l.r, l.c - 1}
	possible := []location{n, s, e, w}
	newLocations := []location{}
	for _, p := range possible {
		isValid := p.isValid(g, prev)
		if isValid {
			newLocations = append(newLocations, p)
		}
	}
	return newLocations
}
