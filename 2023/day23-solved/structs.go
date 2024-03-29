package day23

import "fmt"

type point struct {
	row int
	col int
}

type intersection struct {
	location  point
	adjacents map[point]int
}

type trip struct {
	location intersection
	visited  map[point]bool
	path     []point
	length   int
}

// part1 only
type path struct {
	last    point
	history map[point]bool
}

func (i intersection) getAdjacents(grid []string, end point) map[point]int {
	adjacents := make(map[point]int)
	points := i.location.getValidAdjacents(grid)
	length := 0
	for {
		if len(points) == 0 {
			break
		}
		next := points[0]
		points = points[1:]
		length++

		if next == end {
			adjacents[next] = length
			continue
		}

		grid = markVisit(grid, next)
		nextPoints := next.getValidAdjacents(grid)
		switch len(nextPoints) {
		case 0:
			length = 0
			continue
		case 1:
			points = append(nextPoints, points...)
		default:
			adjacents[next] = length
			length = 0
		}
	}
	return adjacents
}

func (t trip) isFinished(grid []string) bool {
	return t.location.location.isFinished(grid)
}

func (t trip) getNewTrips(grid []string, intersections map[point]intersection) []trip {
	var newTrips []trip
	for nextP, length := range t.location.adjacents {
		if t.visited[nextP] {
			continue
		}
		newVisited := make(map[point]bool)
		for p, _ := range t.visited {
			newVisited[p] = true
		}
		newPath := make([]point, len(t.path))
		copy(newPath, t.path)
		newVisited[nextP] = true
		nextIntersection := intersections[nextP]
		newTrip := trip{nextIntersection, newVisited, append(newPath, nextP), t.length + length}
		newTrips = append(newTrips, newTrip)
	}
	return newTrips
}

func (t trip) toString() string {
	str := fmt.Sprintf("%v: %d (", t.location.location, t.length)
	// for _, p := range t.path {
	// 	str += fmt.Sprintf("%v, ", p)
	// }
	str += ")"
	return str
}

func (p point) getValidAdjacents(grid []string) []point {
	possibles := p.getAdjacent()
	valids := []point{}
	for _, possible := range possibles {
		if possible.isValid(grid, p, true) {
			valids = append(valids, possible)
		}
	}
	return valids
}

func markVisit(grid []string, p point) []string {
	return updateGrid(grid, p, '+')
}

func updateGrid(grid []string, p point, r rune) []string {
	if !p.isValid(grid, p, true) {
		return grid
	}
	row := []rune(grid[p.row])
	row[p.col] = r
	grid[p.row] = string(row)
	return grid
}

func (p point) isValid(grid []string, last point, isPart2 bool) bool {
	return p.row >= 0 &&
		p.row < len(grid) &&
		p.col >= 0 &&
		p.col < len(grid[0]) &&
		grid[p.row][p.col] != '#' &&
		grid[p.row][p.col] != '+' &&
		(isPart2 || p.isDownhill(grid, last))
}

func (p point) isDownhill(grid []string, last point) bool {
	switch grid[p.row][p.col] {
	case '^':
		return last.row > p.row
	case '>':
		return last.col < p.col
	case 'v':
		return last.row < p.row
	case '<':
		return last.col > p.col
	default:
		return true
	}
}

func (p point) getAdjacent() []point {
	return []point{
		{p.row - 1, p.col},
		{p.row, p.col + 1},
		{p.row + 1, p.col},
		{p.row, p.col - 1},
	}
}

func (p path) length() int {
	return len(p.history)
}

func (p path) isValid(grid []string, next point, isPart2 bool) bool {
	return next.isValid(grid, p.last, isPart2) && !p.history[next]
}

func (p path) getNewPaths(grid []string, isPart2 bool) []path {
	var newPaths []path
	for _, next := range p.last.getAdjacent() {
		if p.isValid(grid, next, isPart2) {
			newHistory := make(map[point]bool)
			for k, v := range p.history {
				newHistory[k] = v
			}
			newHistory[next] = true
			newPaths = append(newPaths, path{next, newHistory})
		}
	}
	return newPaths
}

func (p path) isFinished(grid []string) bool {
	return p.last.row == len(grid)-1
}

func (p point) isFinished(grid []string) bool {
	return p.row == len(grid)-1
}
