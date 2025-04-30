package day

import "fmt"

type point struct {
	row int
	col int
}

func main(maze []string, target int) int {
	baseRoute := mapBaseRoute(maze)
	return countCheats(maze, baseRoute, target)
}

func countCheats(maze []string, baseRoute map[point]int, target int) int {
	cheats := 0
	for r, row := range maze {
		for c, cell := range row {
			if cell == '#' {
				continue
			}
			start := point{row: r, col: c}
			baseStartSteps, ok := baseRoute[start]
			if !ok {
				continue
			}

			up := point{row: r - 2, col: c}
			dn := point{row: r + 2, col: c}
			lt := point{row: r, col: c - 2}
			rt := point{row: r, col: c + 2}
			for _, end := range []point{up, dn, lt, rt} {
				baseEndSteps, ok := baseRoute[end]
				if !ok {
					continue
				}
				if !ok {
					panic(fmt.Sprintf("base route not found for start point %v", end))
				}
				stepsAfterCheat := baseStartSteps + 2
				savings := baseEndSteps - stepsAfterCheat
				if savings >= target {
					cheats++
				}
			}
		}
	}
	return cheats
}
