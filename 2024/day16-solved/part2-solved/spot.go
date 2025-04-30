package day

type spot struct {
	row int
	col int
	d   int
}

const (
	up = iota
	dn
	lt
	rt
	none
)

func (s spot) lookAround() []spot {
	return []spot{s.look(up), s.look(dn), s.look(lt), s.look(rt)}
}

func (s spot) look(d int) spot {
	switch d {
	case up:
		return spot{s.row - 1, s.col, d}
	case dn:
		return spot{s.row + 1, s.col, d}
	case lt:
		return spot{s.row, s.col - 1, d}
	case rt:
		return spot{s.row, s.col + 1, d}
	default:
		panic("Invalid direction")
	}
}

func (s spot) isOnBestPath(gridOg []string, best int) bool {
	grid := make([]string, len(gridOg))
	m := maze{spots: []spot{s}, grid: grid, isAtS: false, isAtE: false}

	copy(m.grid, gridOg)
	hasBestTurns := m.hasBestTurns(best / 1000)
	if !hasBestTurns {
		return false
	}

	copy(m.grid, gridOg)
	m.isAtS, m.isAtE = false, false
	m.toS, m.toE = 0, 0
	hasBestSteps := m.hasBestSteps(best % 1000)

	return hasBestTurns && hasBestSteps
}
