package day17

import (
	"fmt"
	"goAoc2023/utils"
)

type direction string

const (
	n direction = "^"
	w direction = "<"
	s direction = "v"
	e direction = ">"
)

type path struct {
	history  string
	heatLoss int
}

func runA_too_slow(grid []string) int {
	debug := true
	// l := point{1, 1}
	// p := path{"v>v>^<", 29}
	// pMap := map[point][]path{}
	// pMap[l] = []path{p}
	// testPaths := filterOutLoops(pMap)
	// fmt.Println(testPaths)
	paths := map[point][]path{}
	paths[point{0, 0}] = []path{{
		history:  "",
		heatLoss: 0,
	}}
	finish := point{len(grid) - 1, len(grid[0]) - 1}
	log(fmt.Sprintf("finish: %v\n", finish), debug)
	count := 0
	for {
		if count > 30 {
			return paths[finish][0].heatLoss
		}
		paths = updatePaths(paths, grid)
		count++

		log(fmt.Sprintf("\n"), debug)
		log(fmt.Sprintf("count: %d\n", count), debug)
		logPaths(paths, debug)
	}
}

func updatePaths(paths map[point][]path, grid []string) map[point][]path {
	debug := false
	newPaths := map[point][]path{}
	for l, ps := range paths {
		copy(newPaths[l], ps)
		for _, p := range ps {
			log(fmt.Sprintf("\n  updating from %v: %v\n", l, p), debug)
			newPs := getValidNewPaths(l, p, grid)
			for newL, newPs := range newPs {
				log(fmt.Sprintf("    newPaths %v: %v\n", newL, newPs), debug)
				newPaths[newL] = append(newPaths[newL], newPs...)
			}
		}
	}
	return filter(newPaths)
}

func getValidNewPaths(l point, p path, grid []string) map[point][]path {
	newPaths := map[point][]path{}
	lastDirection := ""
	if len(p.history) > 0 {
		lastDirection = p.history[len(p.history)-1:]
	}
	repeatMoves := getRecentRepeatedMoves(p.history)
	if lastDirection != "v" &&
		l.row > 0 &&
		(lastDirection != "^" || repeatMoves < 3) {
		newLocation := point{l.row - 1, l.col}
		newPaths[newLocation] = append(newPaths[newLocation], getNewPath(p, newLocation, n, grid))
	}
	if lastDirection != ">" &&
		l.col > 0 &&
		(lastDirection != "<" || repeatMoves < 3) {
		newLocation := point{l.row, l.col - 1}
		newPaths[newLocation] = append(newPaths[newLocation], getNewPath(p, newLocation, w, grid))
	}
	if lastDirection != "^" &&
		l.row < len(grid)-1 &&
		(lastDirection != "v" || repeatMoves < 3) {
		newLocation := point{l.row + 1, l.col}
		newPaths[newLocation] = append(newPaths[newLocation], getNewPath(p, newLocation, s, grid))
	}
	if lastDirection != "<" &&
		l.col < len(grid[0])-1 &&
		(lastDirection != ">" || repeatMoves < 3) {
		newLocation := point{l.row, l.col + 1}
		newPaths[newLocation] = append(newPaths[newLocation], getNewPath(p, newLocation, e, grid))
	}
	return newPaths
}

func getNewPath(oldPath path, l point, d direction, grid []string) path {
	newHistory := oldPath.history + string(d)
	newPath := path{
		heatLoss: oldPath.heatLoss + utils.Atoi(string(grid[l.row][l.col])),
		history:  newHistory,
	}
	return newPath
}

func filter(paths map[point][]path) map[point][]path {
	noDuplicates := filterOutDuplicates(paths)
	return filterOutLoops(noDuplicates)
}

func filterOutDuplicates(paths map[point][]path) map[point][]path {
	filtered := map[point][]path{}
	foundMap := map[string]int{}
	for l, ps := range paths {
		for _, p := range ps {
			if foundMap[p.history] == 1 {
				continue
			}

			foundMap[p.history] = 1
			filtered[l] = append(filtered[l], p)
		}
	}
	return filtered
}

func filterOutLoops(paths map[point][]path) map[point][]path {
	debug := false
	log(fmt.Sprintf("\n  filtering out loops from paths: "), debug)
	filtered := map[point][]path{}
	for l, ps := range paths {
		for _, p := range ps {
			log(fmt.Sprintf("\n  l: %v, %v\n", l, p), debug)
			oldL := point{0, 0}
			for i, step := range p.history {
				log(fmt.Sprintf("    oldL: %v\n", oldL), debug)
				if oldL == l || i == len(p.history)-1 {
					log(fmt.Sprintf("    breaking, looped? %v\n", oldL == l), debug)
					break
				}

				switch step {
				case '^':
					oldL.row -= 1
				case '<':
					oldL.col -= 1
				case 'v':
					oldL.row += 1
				case '>':
					oldL.col += 1
				}
				log(fmt.Sprintf("  step: %c, new oldL: %v\n", step, oldL), debug)
			}
			if oldL == l {
				break
			}

			filtered[l] = append(filtered[l], p)
		}
	}
	return filtered
}

func getLowestFinishedHeatLoss(paths map[string]path) int {
	lowest := -1
	for _, p := range paths {
		if lowest < 0 || p.heatLoss < lowest {
			lowest = p.heatLoss
		}
	}
	return lowest
}

func logPaths(paths map[point][]path, shouldLog bool) {
	if shouldLog {
		for l, ps := range paths {
			for _, p := range ps {
				fmt.Printf(" %v: %v\n", l, p)
			}

		}
	}
}
