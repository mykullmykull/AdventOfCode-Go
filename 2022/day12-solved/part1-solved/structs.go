package day

import (
	"fmt"
)

type grid struct {
	rows      []string
	steps     int
	edgeSpots []spot
}

type spot struct {
	elevation rune
	row       int
	col       int
}

func (g grid) getSpotByRune(r rune) spot {
	for x, row := range g.rows {
		for y, elev := range row {
			if elev == r {
				s := spot{
					row: x,
					col: y,
				}
				// getElevation to currect for 'S' or 'E'
				newElev, err := g.getElevation(s)
				if err != nil {
					panic(fmt.Sprintf("failed to get elev for %v, %v", r, err))
				}
				s.elevation = newElev
				return s
			}
		}
	}
	panic(fmt.Sprintf("No spot found with %v", r))
}

func (g grid) getElevation(s spot) (rune, error) {
	if s.row < 0 || s.row >= len(g.rows) ||
		s.col < 0 || s.col >= len(g.rows[0]) {
		return '-', fmt.Errorf("outside of grid")
	}
	r := rune(g.rows[s.row][s.col])

	if r == 'S' {
		r = 'a'
	}

	return r, nil
}

func (g grid) advanceOneStep() grid {
	newEdges := []spot{}
	for _, s := range g.edgeSpots {
		g = g.markSpot(s, ".")
		up := g.getSpot(s.row-1, s.col)
		dn := g.getSpot(s.row+1, s.col)
		lt := g.getSpot(s.row, s.col-1)
		rt := g.getSpot(s.row, s.col+1)
		possibleNewEdges := []spot{up, dn, lt, rt}
		for _, e := range possibleNewEdges {
			if e.elevation == 'E' && e.isStepableFrom(s) {
				newEdges = []spot{e}
				break
			}
			if byte(e.elevation) < 97 || // 'a'
				byte(e.elevation) > 122 || // 'z'
				!e.isStepableFrom(s) {
				continue
			}
			g = g.markSpot(e, "0")
			newEdges = append(newEdges, e)
		}
	}
	g.edgeSpots = newEdges
	g.steps++
	return g
}

func (g grid) markSpot(s spot, str string) grid {
	newRow := g.rows[s.row]
	newRow = newRow[:s.col] + str + newRow[s.col+1:]
	g.rows[s.row] = newRow
	return g
}

func (g grid) getSpot(r int, c int) spot {
	s := spot{row: r, col: c}
	elev, err := g.getElevation(s)
	if err == nil {
		s.elevation = elev
	}
	return s
}

func (s spot) isStepableFrom(other spot) bool {
	if s.elevation == 'E' {
		return other.elevation == 'z'
	}
	diff := int(byte(s.elevation)) - int(byte(other.elevation))
	return diff <= 1
}

func (g grid) printGrid() {
	for _, r := range g.rows {
		println(r)
	}
	println()
}
