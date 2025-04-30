package day

import "fmt"

type point struct {
	row, col, alt int
}

func main(grid []string) int {
	totalScore := 0
	trails := getTrailHeads(grid)
	for _, trail := range trails {
		score := climb(trail, grid)
		totalScore += score
	}
	return totalScore
}

func getTrailHeads(grid []string) []point {
	trails := []point{}
	for r, row := range grid {
		for c, char := range row {
			if char == '0' {
				trails = append(trails, point{r, c, 0})
			}
		}
	}
	return trails
}

func climb(trail point, grid []string) int {
	trails := []point{trail}
	for x := 0; x < 9; x++ {
		trails = climbOneStep(trails, grid)
	}
	return len(trails)
}

func climbOneStep(trails []point, grid []string) []point {
	newTrails := []point{}
	newAlt := trails[0].alt + 1
	for _, trail := range trails {
		up := point{trail.row - 1, trail.col, newAlt}
		rt := point{trail.row, trail.col + 1, newAlt}
		dn := point{trail.row + 1, trail.col, newAlt}
		lt := point{trail.row, trail.col - 1, newAlt}
		for _, p := range []point{up, rt, dn, lt} {
			if p.row < 0 || p.row >= len(grid) || p.col < 0 || p.col >= len(grid[0]) {
				continue
			}
			if fmt.Sprintf("%c", grid[p.row][p.col]) == fmt.Sprintf("%d", p.alt) {
				newTrails = append(newTrails, p)
			}
		}
	}
	return removeDuplicates(newTrails)
}

func removeDuplicates(trails []point) []point {
	unique := []point{}
	for _, trail := range trails {
		isUnique := true
		for _, u := range unique {
			if trail.row == u.row && trail.col == u.col {
				isUnique = false
				break
			}
		}
		if isUnique {
			unique = append(unique, trail)
		}
	}
	return unique
}

func printGrid(grid []string, trails []point) {
	gridCopy := make([]string, len(grid))
	copy(gridCopy, grid)
	for _, trail := range trails {
		gridCopy[trail.row] = gridCopy[trail.row][:trail.col] + "#" + gridCopy[trail.row][trail.col+1:]
	}
	for _, row := range gridCopy {
		println(row)
	}
}
