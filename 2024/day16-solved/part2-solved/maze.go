package day

type maze struct {
	spots []spot
	grid  []string
	turns int
	steps int
	isAtS bool
	toS   int
	isAtE bool
	toE   int
}

func (m maze) removeDeadEnds() maze {
	hasChanged := true
	for hasChanged {
		hasChanged = false
		for row := range m.grid {
			for col := range m.grid[row] {
				s := spot{row, col, none}
				m.spots = []spot{s}
				if m.isDeadEnd() {
					m = m.updateGrid(s, "#")
					hasChanged = true
				}
			}
		}
	}
	return m
}

func (m maze) isDeadEnd() bool {
	if m.getVal(m.spots[0]) != '.' {
		return false
	}
	paths := len(m.getAdjacents(m.spots[0], '.'))
	starts := len(m.getAdjacents(m.spots[0], 'S'))
	ends := len(m.getAdjacents(m.spots[0], 'E'))
	hasBestTurns := len(m.getAdjacents(m.spots[0], 'T'))
	return paths+starts+ends+hasBestTurns <= 1
}

func (m maze) getVal(s spot) byte {
	return m.grid[s.row][s.col]
}

func (m maze) getAdjacents(s spot, b byte) []spot {
	var adjacents []spot
	for _, newSpot := range s.lookAround() {
		if m.getVal(newSpot) == b {
			adjacents = append(adjacents, newSpot)
		}
	}
	return adjacents
}

func (m maze) updateGrid(s spot, val string) maze {
	m.grid[s.row] = m.grid[s.row][:s.col] + val + m.grid[s.row][s.col+1:]
	return m
}

func (m maze) getNextSpots() maze {
	newSpots := []spot{}
	s := m.spots[0]
	for d := range []int{up, dn, lt, rt} {
		next := s.look(d)
		if m.getVal(next) == 'S' {
			m.isAtS = true
			m.toS = m.turns
		}
		if m.getVal(next) == 'E' {
			m.isAtE = true
			m.toE = m.turns
		}
		if m.getVal(next) == '.' {
			s.d = d
			newSpots = append(newSpots, s)
		}
	}
	m.spots = newSpots
	return m
}

func (m maze) isStillTurning(best int) bool {
	if !m.isAtS && !m.isAtE {
		return m.turns*2 <= best
	}
	if !m.isAtE {
		return m.turns+m.toS <= best
	}
	if !m.isAtS {
		return m.turns+m.toE <= best
	}
	return m.turns+m.toS+m.toE <= best
}

func (m maze) growStraight() maze {
	for _, s := range m.spots {
		for {
			newSpot := s.look(s.d)
			if m.getVal(newSpot) == '#' {
				break
			}
			if !m.isAtS && m.getVal(newSpot) == 'S' {
				m.isAtS = true
				m.toS = m.turns
				if s.d != lt {
					m.toS++
				}
				break
			}
			if !m.isAtE && m.getVal(newSpot) == 'E' {
				m.isAtE = true
				m.toE = m.turns
				break
			}
			m = m.updateGrid(newSpot, "O")
			s = newSpot
		}
	}
	return m
}

func (m maze) turn() maze {
	m.turns++
	newSpots := []spot{}
	for row := range m.grid {
		for col := range m.grid[row] {
			s := spot{row, col, none}
			if m.getVal(s) != 'O' {
				continue
			}
			if len(m.getAdjacents(s, '.')) == 0 {
				continue
			}
			for _, d := range []int{up, dn, lt, rt} {
				newSpot := s.look(d)
				if m.getVal(newSpot) == '.' {
					s.d = d
					newSpots = append(newSpots, s)
				}
			}
		}
	}
	m.spots = newSpots
	return m
}

func (m maze) countXs() int {
	count := 0
	for row := range m.grid {
		for col := range m.grid[row] {
			s := spot{row, col, none}
			if m.getVal(s) == 'X' {
				count++
			}
		}
	}
	return count
}

func (m maze) printGrid() {
	for _, row := range m.grid {
		println(row)
	}
	println()
}

func (m maze) isCorner(s spot) bool {
	m.spots = []spot{s}
	m = m.getNextSpots()
	if len(m.spots) != 2 {
		return false
	}
	switch m.spots[0].d {
	case up:
		return m.spots[1].d != dn
	case dn:
		return m.spots[1].d != up
	case lt:
		return m.spots[1].d != rt
	case rt:
		return m.spots[1].d != lt
	default:
		panic("Invalid direction")
	}
}

func (m maze) growToBestTurns(best int) maze {
	if m.getVal(m.spots[0]) != '.' {
		return m
	}
	m = m.updateGrid(m.spots[0], "O")
	m = m.getNextSpots()
	for m.isStillTurning(best) {
		m = m.growStraight()
		m = m.turn()
	}
	return m
}

func (m maze) hasBestTurns(best int) bool {
	return m.isAtE && m.isAtS && m.toS+m.toE <= best
}

func (m maze) hasBestSteps(best int) bool {
	if m.getVal(m.spots[0]) != '.' {
		return false
	}
	m = m.updateGrid(m.spots[0], "O")
	for m.isStillSpreading(best) {
		m = m.spread()
		m.steps++
	}
	return m.isAtE && m.isAtS && m.toS+m.toE <= best
}

func (m maze) isStillSpreading(best int) bool {
	return (!m.isAtE || !m.isAtS) && m.steps <= best
}

func (m maze) spread() maze {
	newGrid := make([]string, len(m.grid))
	newM := maze{grid: newGrid}
	copy(newM.grid, m.grid)
	for row := range m.grid {
		for col := range m.grid[row] {
			s := spot{row, col, none}
			if m.getVal(s) != '.' && m.getVal(s) != 'S' && m.getVal(s) != 'E' {
				continue
			}
			if len(m.getAdjacents(s, 'O')) > 0 {
				if !m.isAtS && m.getVal(s) == 'S' {
					m.isAtS = true
					m.toS = m.steps + 1
					continue
				}
				if !m.isAtE && m.getVal(s) == 'E' {
					m.isAtE = true
					m.toE = m.steps + 1
					continue
				}
				newM = newM.updateGrid(s, "O")
			}
		}
	}
	copy(m.grid, newM.grid)
	return m
}

func (m maze) filterOutBackTrackers() maze {
	hasChanged := true
	for hasChanged {
		hasChanged = false
		for row := range m.grid {
			for col := range m.grid[row] {
				s := spot{row, col, none}
				if m.getVal(s) != 'X' {
					continue
				}
				if len(m.getAdjacents(s, 'S'))+len(m.getAdjacents(s, 'E')) > 0 {
					continue
				}
				if len(m.getAdjacents(s, 'X')) < 2 {
					m = m.updateGrid(s, ".")
					hasChanged = true
				}
			}
		}
	}
	return m
}

func (m maze) markReachable(other maze) maze {
	for row := range m.grid {
		for col := range m.grid[row] {
			s := spot{row, col, none}
			if other.getVal(s) == 'O' {
				m.updateGrid(s, "R")
			}
		}
	}
	return m
}

func (m maze) markUnreachable() maze {
	for row := range m.grid {
		for col := range m.grid[row] {
			s := spot{row, col, none}
			if m.getVal(s) == '.' {
				m.updateGrid(s, "#")
			}
		}
	}
	for row := range m.grid {
		for col := range m.grid[row] {
			s := spot{row, col, none}
			if m.getVal(s) == 'T' {
				m.updateGrid(s, ".")
			}
		}
	}
	return m
}
