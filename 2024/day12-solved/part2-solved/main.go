package day

import "fmt"

type farm struct {
	grid       []string
	fenceCosts int
	plant      string
	area       int
	perimeter  int
}

type point struct {
	row int
	col int
}

const (
	no   = "#"
	tp   = "-"
	rt   = "]"
	bm   = "_"
	lt   = "["
	tr   = "1"
	tb   = "2"
	tl   = "3"
	trb  = "4"
	tbl  = "5"
	trl  = "6"
	trbl = "7"
	rb   = "8"
	rl   = "9"
	rbl  = "0"
	bl   = "!"
)

func main(grid []string) int {
	f := farm{grid: grid}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if f.grid[r][c] == '.' {
				continue
			}

			f.plant = string(f.grid[r][c])
			println("\nfound new plant", f.plant, "at", r, c)
			f = f.addFenceCost(point{r, c})
			println("area", f.area, "perimeter", f.perimeter, "fenceCost", f.area*f.perimeter)
			f.printGrid()

			f.plant = ""
			f.area = 0
			f.perimeter = 0
		}
	}
	return f.fenceCosts
}

func (f farm) addFenceCost(p point) farm {
	f = f.getAreaAndPerimeter(p)
	fenceCost := f.perimeter * f.area
	f.fenceCosts += fenceCost
	return f
}

func (f farm) getAreaAndPerimeter(p point) farm {
	f = f.markPointsToCount(p)
	f = f.getArea(p)
	f = f.getPerimeter(p)
	f = f.markCounted(p, ".")
	return f
}

func (f farm) markPointsToCount(p point) farm {
	pointsToCount := []point{p}
	for len(pointsToCount) > 0 {
		p := pointsToCount[0]
		pointsToCount = pointsToCount[1:]
		if !f.isLetter(p) {
			continue
		}
		pointsToCount = append(pointsToCount, f.getValidNeighbors(p, f.plant)...)
		sidesSymbol := f.getSidesSymbol(p)
		f = f.mark(p, sidesSymbol)
	}
	return f
}

func (f farm) isLetter(p point) bool {
	return p.row >= 0 &&
		p.row < len(f.grid) &&
		p.col >= 0 &&
		p.col < len(f.grid[0]) &&
		f.grid[p.row][p.col] >= 'A' && f.grid[p.row][p.col] <= 'Z'
}

func (f farm) getValidNeighbors(p point, plant string) []point {
	points := []point{}
	n := point{p.row - 1, p.col}
	s := point{p.row + 1, p.col}
	e := point{p.row, p.col + 1}
	w := point{p.row, p.col - 1}
	for _, p := range []point{n, s, e, w} {
		if f.isValid(p, plant) {
			points = append(points, p)
		}
	}
	return points
}

func (f farm) isValid(p point, plant string) bool {
	return p.row >= 0 &&
		p.row < len(f.grid) &&
		p.col >= 0 &&
		p.col < len(f.grid[0]) &&
		string(f.grid[p.row][p.col]) == plant
}

func (f farm) getSidesSymbol(p point) string {
	p2 := point{p.row - 1, p.col}
	needsTpWall := p2.row < 0 || f.get(p2) == "." || (f.get(p2) != f.plant && f.isLetter(p2))

	p2 = point{p.row + 1, p.col}
	needsBmWall := p2.row >= len(f.grid) || f.get(p2) == "." || (f.get(p2) != f.plant && f.isLetter(p2))

	p2 = point{p.row, p.col + 1}
	needsRtWall := p2.col >= len(f.grid[0]) || f.get(p2) == "." || (f.get(p2) != f.plant && f.isLetter(p2))

	p2 = point{p.row, p.col - 1}
	needsLtWall := p2.col < 0 || f.get(p2) == "." || (f.get(p2) != f.plant && f.isLetter(p2))

	switch [4]bool{needsTpWall, needsRtWall, needsBmWall, needsLtWall} {
	case [4]bool{true, true, true, true}:
		return trbl
	case [4]bool{true, true, true, false}:
		return trb
	case [4]bool{true, true, false, true}:
		return trl
	case [4]bool{true, true, false, false}:
		return tr
	case [4]bool{true, false, true, true}:
		return tbl
	case [4]bool{true, false, true, false}:
		return tb
	case [4]bool{true, false, false, true}:
		return tl
	case [4]bool{true, false, false, false}:
		return tp
	case [4]bool{false, true, true, true}:
		return rbl
	case [4]bool{false, true, true, false}:
		return rb
	case [4]bool{false, true, false, true}:
		return rl
	case [4]bool{false, true, false, false}:
		return rt
	case [4]bool{false, false, true, true}:
		return bl
	case [4]bool{false, false, true, false}:
		return bm
	case [4]bool{false, false, false, true}:
		return lt
	case [4]bool{false, false, false, false}:
		return no
	default:
		panic(fmt.Sprintf("invalid sides for point %d, %d", p.row, p.col))
	}
}

func (f farm) get(p point) string {
	if p.row < 0 ||
		p.row >= len(f.grid) ||
		p.col < 0 ||
		p.col >= len(f.grid[0]) {
		return ""
	}
	return string(f.grid[p.row][p.col])
}

func (f farm) mark(p point, s string) farm {
	f.grid[p.row] = f.grid[p.row][:p.col] + s + f.grid[p.row][p.col+1:]
	return f
}

func (f farm) getArea(p point) farm {
	for r := p.row; r < len(f.grid); r++ {
		for c := 0; c < len(f.grid[0]); c++ {
			if f.get(point{r, c}) == "." {
				continue
			}
			if !f.isLetter(point{r, c}) {
				f.area++
			}
		}
	}
	return f
}

func (f farm) getPerimeter(p point) farm {
	for r := p.row; r < len(f.grid); r++ {
		for c := 0; c < len(f.grid[0]); c++ {
			value := f.get(point{r, c})
			if f.isLetter(point{r, c}) || value == no {
				continue
			}
			if hasTpWall(value) && !hasTpWall(f.get(point{r, c - 1})) {
				f.perimeter++
			}
			if hasRtWall(value) && !hasRtWall(f.get(point{r - 1, c})) {
				f.perimeter++
			}
			if hasBmWall(value) && !hasBmWall(f.get(point{r, c - 1})) {
				f.perimeter++
			}
			if hasLtWall(value) && !hasLtWall(f.get(point{r - 1, c})) {
				f.perimeter++
			}
		}
	}
	return f
}

func hasTpWall(s string) bool {
	return s == tp || s == tr || s == tb || s == tl || s == trb || s == tbl || s == trl || s == trbl
}

func hasRtWall(s string) bool {
	return s == rt || s == tr || s == rb || s == rl || s == trb || s == trl || s == rbl || s == trbl
}

func hasBmWall(s string) bool {
	return s == bm || s == tb || s == rb || s == bl || s == trb || s == tbl || s == rbl || s == trbl
}

func hasLtWall(s string) bool {
	return s == lt || s == tl || s == bl || s == rl || s == trl || s == tbl || s == rbl || s == trbl
}

func (f farm) markCounted(p point, s string) farm {
	for r := p.row; r < len(f.grid); r++ {
		for c := 0; c < len(f.grid[0]); c++ {
			if !f.isLetter(point{r, c}) {
				f = f.mark(point{r, c}, s)
			}
		}
	}
	return f
}

func (f farm) printGrid() {
	for _, row := range f.grid {
		println(row)
	}
}
