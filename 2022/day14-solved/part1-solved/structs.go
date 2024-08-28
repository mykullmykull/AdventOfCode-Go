package day

import "math"

type grid struct {
	rows      []string
	origin    spot
	sandCount int
}

type spot struct {
	row int
	col int
}

type rockLine struct {
	start spot
	end   spot
}

func (g grid) addRocks(rockLines []rockLine) grid {
	for _, rl := range rockLines {
		s := spot{col: rl.start.col, row: rl.start.row}
		colChange := rl.end.col - rl.start.col
		if colChange != 0 {
			colChange = colChange / int(math.Abs(float64(colChange)))
		}
		rowChange := rl.end.row - rl.start.row
		if rowChange != 0 {
			rowChange = rowChange / int(math.Abs(float64(rowChange)))
		}
		for {
			g = g.markSpot(s, "#")
			if s.col == rl.end.col && s.row == rl.end.row {
				break
			}
			s.col += colChange
			s.row += rowChange
		}
	}
	return g
}

func (g grid) markSpot(s spot, mark string) grid {
	row := g.rows[s.row]
	row = row[:s.col] + mark + row[s.col+1:]
	g.rows[s.row] = row
	return g
}

func (g grid) printGrid(s spot) {
	copyRows := make([]string, len(g.rows))
	copy(copyRows, g.rows)
	copyG := grid{rows: copyRows}
	copyG.markSpot(s, "+")
	for _, row := range copyG.rows {
		println(row[460:])
	}
}

func (g grid) addSandUntilVoid() grid {
	s := spot{col: g.origin.col, row: g.origin.row}
	for {
		newS := s.fall(g)

		// sand fell into the void
		if newS.row >= len(g.rows) {
			return g
		}

		// sand stopped
		if newS.row == s.row {
			g = g.markSpot(s, "o")
			g.sandCount += 1
			s = spot{col: g.origin.col, row: g.origin.row}
			continue
		}

		s = newS
	}
}

func (s spot) fall(g grid) spot {
	// falls down
	dn := spot{col: s.col, row: s.row + 1}
	if g.isEmpty(dn) {
		return dn
	}

	// falls left
	lt := spot{col: s.col - 1, row: s.row + 1}
	if g.isEmpty(lt) {
		return lt
	}

	// falls right
	rt := spot{col: s.col + 1, row: s.row + 1}
	if g.isEmpty(rt) {
		return rt
	}

	return s
}

func (g grid) isEmpty(s spot) bool {
	return s.row >= len(g.rows) || g.rows[s.row][s.col] == '.'
}
