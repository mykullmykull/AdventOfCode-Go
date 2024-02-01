package day17

import (
	"fmt"
	"goAoc2023/utils"
)

type edge struct {
	location point
	heatLoss int
	history  string
}

func runA_misses_optimizing_each_step(grid []string) int {
	debug := true
	firstPoint := point{0, 0}
	grid = updateGrid(grid, firstPoint, "*")
	edges := []edge{{firstPoint, 0, ""}}
	finish := point{len(grid) - 1, len(grid[0]) - 1}
	for {
		logGrid(grid, debug)
		logEdges(edges, debug)

		finalHeatLoss := getFinalHeatLoss(edges, finish)
		if finalHeatLoss > 0 {
			return finalHeatLoss
		}
		grid, edges = spread(grid, edges, finish)
	}
}

func getFinalHeatLoss(edges []edge, finish point) int {
	for _, e := range edges {
		if e.location.row == finish.row && e.location.col == finish.col {
			return e.heatLoss
		}
	}
	return -1
}

func spread(grid []string, edges []edge, finish point) ([]string, []edge) {
	newEdge := getNewEdgeFromLowestValidBoarder(edges, grid)
	edges = append(edges, newEdge)
	grid, edges = updateAfterSpread(grid, newEdge.location, edges, finish)
	return grid, edges
}

func getNewEdgeFromLowestValidBoarder(edges []edge, grid []string) edge {
	debug := false
	log("  looking for lowest unmarked point", debug)
	lowestHeat := 0
	lowestPoint := point{0, 0}
	bestAdjacentEdge := edge{}
	newDirection := ""
	for r, row := range grid {
		for c, char := range row {
			log(fmt.Sprintf("    considering %d, %d", r, c), debug)
			if isMarked(char) {
				log("    already marked", debug)
				continue
			}

			if !isOnEdge(r, c, grid, false) {
				log("    not on the border", debug)
				continue
			}
			heatLoss := utils.Atoi(string(char))
			if lowestHeat == 0 || heatLoss < lowestHeat {
				p := point{r, c}
				bestEdgeForP, direction := getBestAdjacentEdge(edges, p)
				log(fmt.Sprintf("    bestAdjacentEdge: %v, newDirection: %s", bestAdjacentEdge, newDirection), debug)
				if bestEdgeForP.heatLoss == -1 {
					continue
				}
				log(fmt.Sprintf("    heatLoss: %d, lowestHeat: %d", heatLoss, lowestHeat), debug)
				lowestHeat = heatLoss
				lowestPoint = p
				bestAdjacentEdge = bestEdgeForP
				newDirection = direction
			}
		}
	}
	log(fmt.Sprintf("    final lowest heat: %d at point: %v", lowestHeat, lowestPoint), debug)
	return edge{lowestPoint, bestAdjacentEdge.heatLoss + lowestHeat, bestAdjacentEdge.history + newDirection}
}

func isMarked(char rune) bool {
	return char == '#' || char == '*'
}

func isOnEdge(r int, c int, grid []string, inside bool) bool {
	if (inside && !isMarked(getRuneFromGrid(grid, point{r, c}))) ||
		(!inside && isMarked(getRuneFromGrid(grid, point{r, c}))) {
		return false
	}

	up := point{r - 1, c}
	dn := point{r + 1, c}
	lt := point{r, c - 1}
	rt := point{r, c + 1}

	for _, p := range []point{up, dn, lt, rt} {
		if p.row < 0 || p.row >= len(grid) ||
			p.col < 0 || p.col >= len(grid[0]) {
			continue
		}
		if (inside && !isMarked(getRuneFromGrid(grid, p))) ||
			(!inside && isMarked(getRuneFromGrid(grid, p))) {
			return true
		}
	}
	return false
}

func updateAfterSpread(grid []string, p point, edges []edge, finish point) ([]string, []edge) {
	newEdges := []edge{}
	grid = removeOldStars(grid)
	grid = updateGrid(grid, p, "*")
	for _, e := range edges {
		if e.location.row == finish.row && e.location.col == finish.col {
			return grid, []edge{e}
		}

		if isOnEdge(e.location.row, e.location.col, grid, true) {
			newEdges = append(newEdges, e)
		}
	}
	return grid, newEdges
}

func getBestAdjacentEdge(edges []edge, p point) (edge, string) {
	debug := false
	bestEdge := edge{point{}, -1, ""}
	newDirection := ""

	up := point{p.row - 1, p.col}
	dn := point{p.row + 1, p.col}
	lt := point{p.row, p.col - 1}
	rt := point{p.row, p.col + 1}

	for _, adjacentP := range []point{up, dn, lt, rt} {
		for _, e := range edges {
			if e.location.row != adjacentP.row || e.location.col != adjacentP.col {
				continue
			}

			fromEdgeToPoint := getDirection(e, p)
			log(fmt.Sprintf("      considering edge: %v to point: %v, with this direction: %s", e, p, fromEdgeToPoint), debug)
			if getRecentRepeatedMoves(e.history+fromEdgeToPoint) > 3 {
				log("      too many repeats", debug)
				continue
			}

			if bestEdge.heatLoss == -1 || e.heatLoss < bestEdge.heatLoss {
				log("      found better edge", debug)
				bestEdge = e
				newDirection = fromEdgeToPoint
			}
		}
	}
	return bestEdge, newDirection
}

func getDirection(e edge, p point) string {
	direction := ""

	switch e.location.row - p.row {
	case -1:
		direction = "v"
	case 1:
		direction = "^"
	}

	switch e.location.col - p.col {
	case -1:
		direction = ">"
	case 1:
		direction = "<"
	}
	return direction
}

func removeOldStars(grid []string) []string {
	for r, row := range grid {
		for c, char := range row {
			if char == '*' {
				grid = updateGrid(grid, point{r, c}, "#")
			}
		}
	}
	return grid
}

func logGrid(grid []string, shouldLog bool) {
	if !shouldLog {
		return
	}

	fmt.Println()
	fmt.Println()
	for _, row := range grid {
		fmt.Println(row)
	}
	return
}

func logEdges(edges []edge, shouldLog bool) {
	if !shouldLog {
		return
	}

	fmt.Println()
	for _, e := range edges {
		fmt.Println(e)
	}
	return
}
