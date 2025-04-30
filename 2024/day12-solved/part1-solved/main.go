package day

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
	f = f.countPoints(p)
	f = f.markPoints(p, ".")
	return f
}

func (f farm) markPointsToCount(p point) farm {
	pointsToCount := []point{p}
	for len(pointsToCount) > 0 {
		p := pointsToCount[0]
		pointsToCount = pointsToCount[1:]
		if f.grid[p.row][p.col] == '#' {
			continue
		}
		pointsToCount = append(pointsToCount, f.getValidNeighbors(p, f.plant)...)
		f = f.mark(p, "#")
	}
	return f
}

func (f farm) countPoints(p point) farm {
	for r := p.row; r < len(f.grid); r++ {
		for c := 0; c < len(f.grid[0]); c++ {
			if f.grid[r][c] == '#' {
				f = f.countPoint(point{r, c})
			}
		}
	}
	return f
}

func (f farm) countPoint(p point) farm {
	f.area++
	edges := 4 - len(f.getValidNeighbors(p, "#"))
	f.perimeter += edges
	return f
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

func (f farm) mark(p point, s string) farm {
	f.grid[p.row] = f.grid[p.row][:p.col] + s + f.grid[p.row][p.col+1:]
	return f
}

func (f farm) markPoints(p point, s string) farm {
	for r := p.row; r < len(f.grid); r++ {
		for c := 0; c < len(f.grid[0]); c++ {
			if f.grid[r][c] == '#' {
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
