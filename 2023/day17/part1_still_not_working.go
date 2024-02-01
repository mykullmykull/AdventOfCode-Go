package day17

import (
	"fmt"
	"goAoc2023/utils"
)

//STRATEGY
/*
	Total min heat loss and best path to get to each block, then take the heat loss of the final block
	Working in diagonals from start to finish
	Identify the values for each block by progressing from each of the last diagnoals and spreading like water along the smallest path
*/

func runA_Still_Not_Working(grid []string) int {
	histories := map[point]history{}
	histories[point{0, 0}] = history{0, ""}
	for distance := 1; distance < len(grid[0])+len(grid)-1; distance++ {
		for row := 0; row <= distance; row++ {
			col := distance - row
			p := point{row, col}
			histories[p] = getBestHistory(p, histories, distance-1, grid)
		}
		for row := 0; row <= distance; row++ {
			col := distance - row
			p := point{row, col}
			updateGrid(grid, p, "#")
		}
	}
	finish := point{len(grid) - 1, len(grid[0]) - 1}
	return histories[finish].heat
}

func getBestHistory(targetP point, histories map[point]history, prevDistance int, grid []string) history {
	bestHistory := history{0, ""}
	gridCopy := make([]string, len(grid))
	for row := 0; row <= prevDistance; row++ {
		col := prevDistance - row
		startP := point{row, col}
		startH := histories[startP]
		copy(gridCopy, grid)
		h := spreadToTarget(targetP, startP, startH, gridCopy)
		if bestHistory.heat == 0 || h.heat < bestHistory.heat {
			bestHistory = h
		}
	}
	return bestHistory
}

func spreadToTarget(targetP point, startP point, startH history, grid []string) history {
	newPath := startH.path
	newHeat := startH.heat
	if startP.isUpFrom(targetP) {
		targetHeat := utils.Atoi(string(grid[targetP.row][targetP.col]))
		newPath += "v"
		newHeat += targetHeat
	}

	if startP.isLtFrom(targetP) {
		targetHeat := utils.Atoi(string(grid[targetP.row][targetP.col]))
		newPath += ">"
		newHeat += targetHeat
	}

	newH := history{newHeat, newPath}
	if newH.hasValidPath(targetP) {
		return newH
	}

	histories := map[point]history{}
	histories[startP] = startH
	for {
		// since grid is a slice, updates will carry with it; no need to return it
		histories = spreadToNextSmallest(histories, grid)
		if startP == targetP {
			return histories[startP]
		}
	}
}

// since grid is a slice, updates will carry with it; no need to return it
func spreadToNextSmallest(histories map[point]history, grid []string) map[point]history {
	edges := getEdgesSortedByHeat(grid)
	for _, e := range edges {
		eHeat := utils.Atoi(string(grid[e.row][e.col]))
		adjacentSources, directionsToEdge := getAdjacentSourcesSortedByHeat(e, histories)
		for i, s := range adjacentSources {
			sourceH := histories[s]
			directionToE := directionsToEdge[i]
			h := history{sourceH.heat + eHeat, sourceH.path + directionToE}
			if h.hasValidPath(e) {
				grid = updateGrid(grid, e, "*")
				histories[e] = h
				return histories
			}
		}
	}
	panic(fmt.Sprintf("can't find next smallest"))
}

func getEdgesSortedByHeat(grid []string) []point {
	edges := []point{}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			char := grid[row][col]
			if char == '#' {
				continue
			}

			up := point{row - 1, col}
			dn := point{row + 1, col}
			lt := point{row, col - 1}
			rt := point{row, col + 1}
			for _, p := range []point{up, dn, lt, rt} {
				if isOutsideGrid(p, grid) {
					continue
				}
				if grid[p.row][p.col] == '*' {
					edges = append(edges, p)
					break
				}
			}
		}
	}

	heatSorted := []int{}
	edgesSorted := []point{}
	for _, e := range edges {
		eHeat := getHeat(e, grid)
		if len(heatSorted) == 0 {
			heatSorted = append(heatSorted, eHeat)
			edgesSorted = append(edgesSorted, e)
			continue
		}

		indexOfBigger := -1
		for i, h := range heatSorted {
			if h > eHeat {
				indexOfBigger = i
				break
			}
		}
		heatSorted = append(heatSorted[:indexOfBigger+1], heatSorted[indexOfBigger])
		heatSorted[indexOfBigger] = eHeat

		edgesSorted = append(edgesSorted[:indexOfBigger+1], edgesSorted[indexOfBigger])
		edgesSorted[indexOfBigger] = e
	}
	return edgesSorted
}

func getAdjacentSourcesSortedByHeat(p point, histories map[point]history) ([]point, []string) {

}

//---------------------------- Util Functions ------------------------//
func updateGrid(grid []string, p point, replacement string) []string {
	row := grid[p.row]
	remainder := ""
	if p.col < len(grid[0]) {
		remainder = row[p.col+1:]
	}
	row = row[:p.col] + replacement + remainder
	grid[p.row] = row
	return grid
}

func printGrid(heatLoss int, grid []string) {
	fmt.Println()
	fmt.Println(heatLoss)
	for _, line := range grid {
		fmt.Println(line)
	}
}

func isOutsideGrid(p point, grid []string) bool {
	return p.row < 0 || p.row >= len(grid) ||
		p.col < 0 || p.col >= len(grid[0])
}

func getHeat(p point, grid []string) int {
	return utils.Atoi(string(grid[p.row][p.col]))
}

//---------------------------- Point struct and functions -----------------//
type point struct {
	row int
	col int
}

func (p point) isUpFrom(p2 point) bool {
	return p.col == p2.col && p2.row-p.row == 1
}

func (p point) isLtFrom(p2 point) bool {
	return p.row == p2.row && p2.col-p.col == 1
}

//---------------------------- History struct and functions ---------------//
type history struct {
	heat int
	path string
}

func (h history) hasValidPath(p point) bool {
	if loops(p, h.path) {
		return false
	}
	if getRecentRepeatedMoves(h.path) > 3 {
		return false
	}
	return true
}

//---------------------------- edge struct and functions ---------------//
type spreadPoint struct {
	p    point
	heat int
}
