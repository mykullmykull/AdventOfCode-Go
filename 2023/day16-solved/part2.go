package day16

import "fmt"

func runB(grid []string) int {
	maxEnergized := 0
	for r := -1; r <= len(grid); r++ {
		for c := -1; c <= len(grid[0]); c++ {
			// cant start in the middle of the grid
			if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
				continue
			}

			// can't start in corners
			if (r == -1 && c == -1) ||
				(r == -1 && c == len(grid[0])) ||
				(r == len(grid) && c == -1) ||
				(r == len(grid) && c == len(grid[0])) {
				continue
			}

			gridMap := parseGridMap(grid)
			initialBeam := getInitialBeam(r, c, len(grid), len(grid[0]))
			gridMap, energizedCount := getFinalEnergizedCount(gridMap, []beam{initialBeam})
			if energizedCount > maxEnergized {
				maxEnergized = energizedCount
			}
		}
	}
	return maxEnergized
}

func getInitialBeam(r int, c int, maxR int, maxC int) beam {
	beam := beam{point{r, c}, rt}
	switch beam.location.row {
	case -1:
		beam.direction = dn
	case maxR:
		beam.direction = up
	default:
		switch beam.location.col {
		case -1:
			beam.direction = rt
		case maxC:
			beam.direction = lt
		default:
			panic(fmt.Sprintf("unexpected initial beam location: %v", beam.location))
		}
	}
	return beam
}

func getFinalEnergizedCount(gridMap map[point]gridPoint, beams []beam) (map[point]gridPoint, int) {
	for {
		gridMap, beams = update(gridMap, beams)
		if len(beams) == 0 {
			break
		}
	}
	return gridMap, countEnergizedGridPoints(gridMap)
}
