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
	remainder := ""
	if s.col < len(g.rows[0]) {
		remainder = row[s.col+1:]
	}
	row = row[:s.col] + mark + remainder
	g.rows[s.row] = row
	return g
}

func (g grid) printGrid(s spot) {
	copyRows := make([]string, len(g.rows))
	copy(copyRows, g.rows)
	copyG := grid{rows: copyRows}
	copyG.markSpot(s, "+")
	for _, row := range copyG.rows {
		println(row)
	}
}

func (g grid) addSandUntilFull() grid {
	s := spot{col: g.origin.col, row: g.origin.row}
	newS := spot{}
	for {
		g, newS = g.dropSand(s)

		// sand stopped blocking origin
		if newS.equals(g.origin) {
			g.sandCount += 1
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

func (g grid) dropSand(s spot) (grid, spot) {
	// stops at bottom
	if s.row == len(g.rows)-1 {
		return g, s
	}

	// falls down
	dn := spot{col: s.col, row: s.row + 1}
	if g.isEmpty(dn) {
		return g, dn
	}

	// falls left
	lt := spot{col: s.col - 1, row: s.row + 1}
	if lt.col < 0 {
		g = g.addCol(true)
		return g, lt
	}
	if g.isEmpty(lt) {
		return g, lt
	}

	// falls right
	rt := spot{col: s.col + 1, row: s.row + 1}
	if rt.col >= len(g.rows[0]) {
		g = g.addCol(false)
		return g, rt
	}
	if g.isEmpty(rt) {
		return g, rt
	}

	return g, s
}

func (g grid) isEmpty(s spot) bool {
	return s.row >= len(g.rows) || g.rows[s.row][s.col] == '.'
}

func (s spot) equals(s2 spot) bool {
	return s.row == s2.row && s.col == s2.col
}

func (g grid) addCol(addToLeft bool) grid {
	newRows := make([]string, len(g.rows))
	for x, oldRow := range g.rows {
		if addToLeft {
			newRows[x] = "." + oldRow
			continue
		}
		newRows[x] = oldRow + "."
	}
	copy(g.rows, newRows)
	return g
}
