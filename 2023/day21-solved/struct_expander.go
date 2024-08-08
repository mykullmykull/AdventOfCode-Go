package day21

import "strings"

type expander struct {
	rows []string
	up   bool
	dn   bool
	lt   bool
	rt   bool
}

func (e expander) expandRows() []string {
	if e.up && e.dn {
		e = e.expandUpAndDn()
	}

	if e.lt && e.rt {
		e = e.expandLtAndRt()
	}
	if e.up || e.dn {
		e = e.expandUpOrDn()
	}
	if e.lt || e.rt {
		e = e.expandLtOrRt()
	}
	return e.rows
}

func (e expander) expandUpAndDn() expander {
	newRows := make([]string, len(e.rows)*3)
	for i := 0; i < 3; i++ {
		for j, row := range e.rows {
			row = strings.Replace(row, "S", ".", -1)
			newIndex := i*len(e.rows) + j
			newRows[newIndex] = row
		}
	}
	e.rows = newRows
	e.up = false
	e.dn = false
	return e
}

func (e expander) expandLtAndRt() expander {
	newRows := make([]string, len(e.rows))
	for i, row := range e.rows {
		row = strings.Replace(row, "S", ".", -1)
		newRows[i] = row + row + row
	}
	e.rows = newRows
	e.lt = false
	e.rt = false
	return e
}

func (e expander) expandUpOrDn() expander {
	newRows := make([]string, len(e.rows)*2)
	for i := 0; i < 2; i++ {
		for j, row := range e.rows {
			row = strings.Replace(row, "S", ".", -1)
			newIndex := i*len(e.rows) + j
			newRows[newIndex] = row
		}
	}
	e.rows = newRows
	e.up = false
	e.dn = false
	return e
}

func (e expander) expandLtOrRt() expander {
	newRows := make([]string, len(e.rows))
	for i, row := range e.rows {
		row = strings.Replace(row, "S", ".", -1)
		newRows[i] = row + row
	}
	e.rows = newRows
	e.lt = false
	e.rt = false
	return e
}

func (e expander) adjustEdgeSpots(edgeSpots []spot) []spot {
	newEdgeSpots := []spot{}
	sfSize := len(e.rows)
	sfEnd := sfSize - 1
	for _, s := range edgeSpots {
		newS := spot{s.r, s.c}
		if e.up && s.r != 0 {
			newS.r = s.r + sfSize
		}
		if e.dn && s.r == sfEnd {
			newS.r = s.r + sfSize
		}
		if e.lt && s.c != 0 {
			newS.c = s.c + sfSize
		}
		if e.rt && s.c == sfEnd {
			newS.c = s.c + sfSize
		}
		newEdgeSpots = append(newEdgeSpots, newS)
	}

	return newEdgeSpots
}
