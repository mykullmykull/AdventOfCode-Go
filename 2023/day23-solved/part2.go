package day23

import "fmt"

func runB(grid []string) int {
	start := intersection{location: point{0, 1}, adjacents: map[point]int{}}
	endPoint := point{len(grid) - 1, len(grid[0]) - 2}
	end := intersection{location: endPoint, adjacents: map[point]int{}}
	intersections := map[point]intersection{}
	intersections[start.location] = start
	intersections[end.location] = end
	intersections = getIntersectionsFromGrid(intersections, grid)
	intersections = mapAdjacents(intersections, grid, endPoint)
	return getMaxLength(intersections, grid)
}

func getIntersectionsFromGrid(intersections map[point]intersection, grid []string) map[point]intersection {
	for i, row := range grid {
		for j, char := range row {
			if char == '#' {
				continue
			}
			p := point{i, j}
			if len(p.getValidAdjacents(grid)) > 2 {
				intersections[p] = intersection{p, map[point]int{}}
			}
		}
	}
	return intersections
}

func mapAdjacents(intersections map[point]intersection, gridOriginal []string, end point) map[point]intersection {
	grid := make([]string, len(gridOriginal))
	for p, i := range intersections {
		copy(grid, gridOriginal)
		grid = markVisit(grid, p)
		i.adjacents = i.getAdjacents(grid, end)
		intersections[p] = i
	}
	return intersections
}

func getMaxLength(intersections map[point]intersection, grid []string) int {
	start := point{0, 1}
	t := trip{
		intersections[start],
		map[point]bool{start: true},
		[]point{start},
		0,
	}
	trips := []trip{t}
	maxLength := 0
	for {
		if len(trips) == 0 {
			break
		}
		next := trips[0]
		fmt.Printf("next: %v of trips %d\n", next.toString(), len(trips))
		trips = trips[1:]

		if next.isFinished(grid) {
			if next.length > maxLength {
				maxLength = next.length
			}
			continue
		}

		newTrips := next.getNewTrips(grid, intersections)
		trips = append(trips, newTrips...)
	}
	return maxLength
}

func printIntersections(intersections map[point]intersection, grid []string) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			p := point{r, c}
			if i, ok := intersections[p]; ok {
				println()
				fmt.Printf("%v:\n", p)
				for q, l := range i.adjacents {
					fmt.Printf("  %v: %d\n", q, l)
				}
			}
		}
	}
}
