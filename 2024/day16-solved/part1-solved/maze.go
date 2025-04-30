package day

import "fmt"

type maze struct {
	grid   []string
	scores map[int]int
	racers []racer
	start  spot
	end    spot
}

func (m maze) initialize() maze {
	m.scores = map[int]int{}
	for row := range m.grid {
		for col := range m.grid[row] {
			r := racer{s: spot{row, col}, d: rt, score: 0}
			if r.s.getVal(m.grid) == 'S' {
				m.start = spot{row, col}
				continue
			}
			if r.s.getVal(m.grid) == 'E' {
				m.end = spot{row, col}
				continue
			}
			m = m.removeDeadEnds()
		}
	}
	m.racers = []racer{{s: m.start, d: rt}}
	return m
}

func (m maze) removeDeadEnds() maze {
	hasChanged := true
	for hasChanged {
		hasChanged = false
		for row := range m.grid {
			for col := range m.grid[row] {
				r := racer{s: spot{row, col}, d: rt, score: 0}
				if r.isDeadEnd(m.grid) {
					m = m.updateGrid(r.s, "#")
					hasChanged = true
				}
			}
		}
	}
	return m
}

func (m maze) updateGrid(s spot, val string) maze {
	m.grid[s.row] = m.grid[s.row][:s.col] + val + m.grid[s.row][s.col+1:]
	return m
}

func (m maze) process() maze {
	r := m.racers[0]
	if r.s.getVal(m.grid) == 'E' {
		m.racers = []racer{}
		return m
	}

	m.racers = m.racers[1:]
	adjacents := r.getValidAdjacents(m.grid)
	for _, adj := range adjacents {
		if r.isBackwards(adj) {
			continue
		}
		if found, ok := m.scores[adj.s.toInt()]; !ok || found > r.score {
			m = m.addRacer(adj)
		}
	}
	return m
}

func (m maze) addRacer(r2 racer) maze {
	m.scores[r2.s.toInt()] = r2.score
	if len(m.racers) == 0 {
		m.racers = append(m.racers, r2)
		return m
	}
	inserted := false
	for i, r1 := range m.racers {
		// if r2.getDistanceToEnd(m.end) > r1.getDistanceToEnd(m.end) {
		// 	continue
		// }
		// if r2.getDistanceToEnd(m.end) < r1.getDistanceToEnd(m.end) ||
		if r2.score < r1.score {
			m.racers = append(m.racers[:i], append([]racer{r2}, m.racers[i:]...)...)
			inserted = true
			break
		}
	}
	if !inserted {
		m.racers = append(m.racers, r2)
	}
	return m
}

func (m maze) printGrid() {
	println()
	for _, row := range m.grid {
		fmt.Println(row)
	}
}
