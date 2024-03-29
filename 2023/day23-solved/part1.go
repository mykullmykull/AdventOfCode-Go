package day23

import "fmt"

func runA(grid []string) int {
	start := point{0, 1}
	firstPath := path{start, map[point]bool{start: true}}
	paths := []path{firstPath}
	maxLength := 0
	for {
		if len(paths)%100 == 0 {
			fmt.Printf("paths: %d, maxLength: %d\n", len(paths), maxLength)
		}
		if len(paths) == 0 {
			break
		}
		nextPath := paths[0]
		paths = paths[1:]
		if nextPath.isFinished(grid) && nextPath.length() > maxLength {
			maxLength = nextPath.length() - 1 // don't count the starting point
			continue
		}
		newPaths := nextPath.getNewPaths(grid, false)
		paths = append(paths, newPaths...)
	}
	return maxLength
}

func printGridPt1(grid []string, paths []path) {
	fmt.Print("\033[H\033[2J")
	for i, row := range grid {
		for j, char := range row {
			point := point{i, j}
			isLast := false
			for _, p := range paths {
				if point == p.last {
					isLast = true
					break
				}
			}
			if isLast {
				print("X")
			} else {
				print(string(char))
			}
		}
		println()
	}
	println()
}
