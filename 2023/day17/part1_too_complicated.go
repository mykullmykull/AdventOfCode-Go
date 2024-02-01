package day17

import (
	"fmt"
	"goAoc2023/utils"
)

/*
	Identify the least amount of heat lost and the best path to get to each point in the grid, then return the answer for the final point
	Work along diagonals to start with the easiest and to limit iterations
	For each focussed point (targetP) in question:
		Try to reach targetP from every already finalized point
			Get points (startP) from previous diagonal (distance = current distance - 1)
			Is startP already finalized?
			Consider all valid paths (path) from startP to foccusP
				TODO: How to get next valid path?
				Update if bestHeatLoss hasn't been set yet or heat loss < bestHeatLoss


	To get next valid path:
	If last path was valid then add ^,
		Is path valid?
			doesn't go outside of grid,
			doesn't go inside finalized
			doesn't repeat step too often,
			doesn't loop
		If not valid then retry
		If at targetP then return
	If last path was not valid then consider last attempted move and
		if ^ undo path and point and try >
		if > undo path and point and try v
		if v undo path and point and try <
		if < undo path and point twice and retry
*/

func runA_too_complicated(grid []string) int {
	debug := true
	finalized := map[point]history{}
	finalized[point{0, 0}] = history{0, ""}
	finalized[point{0, 1}] = history{utils.Atoi(string(grid[0][1])), ">"}
	finalized[point{1, 0}] = history{utils.Atoi(string(grid[1][0])), "v"}

	// for distance := 2; distance < len(grid)+len(grid[0])-1; distance++ {
	for distance := 2; distance < 6; distance++ {
		for row := 0; row <= distance; row++ {
			col := distance - row
			p := point{row, col}
			if isOutsideGrid(p, grid) {
				continue
			}
			pHeatLoss := utils.Atoi(string(grid[p.row][p.col]))
			log(fmt.Sprintf("\ntotalling point: %v with heat loss of %d", p, pHeatLoss), debug)
			finalized[p] = getLeastHeatLost(p, finalized, grid, distance-1)
			log(fmt.Sprintf("totalled to %v", finalized[p]), debug)
			if p.equals(point{0, 4}) {
				break
			}
		}
	}

	finish := point{len(grid) - 1, len(grid[0]) - 1}
	log(fmt.Sprintf("final path: %s", finalized[finish].path), true)
	return finalized[finish].heat
}

func getLeastHeatLost(p point, finalized map[point]history, grid []string, finalizedDistance int) history {
	debug := true

	h := history{0, ""}

	for row := 0; row <= finalizedDistance; row++ {
		col := finalizedDistance - row
		p2 := point{row, col}
		log(fmt.Sprintf("\n  starting at %v: %v", p2, finalized[p2]), debug)
		h = updateBestHistory(p2, p, h, finalized, grid, finalizedDistance)
		log(fmt.Sprintf("  best history: %v", h), debug)
	}
	return h
}

func updateBestHistory(p point, targetP point, bestHistory history, finalized map[point]history, grid []string, distance int) history {
	debug := true
	h := finalized[p]
	hasReachedTargetP := false
	lastMove := ""
	for {
		log(fmt.Sprintf("\n    h: %v, bestHistory: %v, looking for next valid history", h, bestHistory), debug)
		p, h = getNextValidHistory(p, targetP, h, hasReachedTargetP, distance, grid, bestHistory.heat, lastMove)
		hasReachedTargetP = true

		if h.heat == 0 || h.path == "" {
			log("    h.heat == 0", debug)
			return bestHistory
		}

		if bestHistory.heat == 0 || h.heat < bestHistory.heat {
			log(fmt.Sprintf("    found a better history: %v", h), debug)
			bestHistory = h
		}

		// rewind to find next valid path
		log("    rewinding to find next valid path", debug)
		lastMove, p, h = rewind(p, h, grid)
	}
}

func getNextValidHistory(
	startP point,
	targetP point,
	h history,
	shouldRewind bool,
	distance int,
	grid []string,
	bestHeat int,
	lastMove string,
) (point, history) {
	debug := true
	p := startP
	for {
		log(fmt.Sprintf("\n      p: %v, h: %v, shouldRewind: %v", p, h, shouldRewind), debug)
		// if current path is valid and we've reached target p return the current point and history
		if (!shouldRewind && p.equals(targetP)) ||
			(shouldRewind && p.equals(startP)) {
			break
		}

		// if path just needs another step, start with ^
		if !shouldRewind {
			// if heat has already exceeded the present best, then rewind and try again
			if bestHeat > 0 && h.heat > bestHeat {
				log("      heat loss was too much, rewinding", debug)
				shouldRewind = true
				continue
			}

			log("      trying ^", debug)
			p = updatePoint(p, "^", false)
			h = updateHistory(p, h, "^", grid)
			if isPathValid(p, h.path, grid, distance) {
				log("      path is valid", debug)
				// if new path is valid then continue adding more steps
				continue
			}
		}

		// if rewinding or if the new ^ is invalid then rewind path and p and try a new move
		if lastMove == "" {
			lastMove, p, h = rewind(p, h, grid)
		}
		log(fmt.Sprintf("      path was invalid, lastMove: %s", lastMove), debug)
		log(fmt.Sprintf("      reset path: %v, reset p: %v", h.path, p), debug)

		nextMove := ""
		switch lastMove {
		case "^":
			nextMove = "<"
		case "<":
			nextMove = "v"
		case "v":
			nextMove = ">"
		case ">":
			// We've exhasted all moves here and need to rewind again
			lastMove = ""
			shouldRewind = true
			continue
		}
		log(fmt.Sprintf("      trying next move: %s", nextMove), debug)

		// try the next move
		p = updatePoint(p, nextMove, false)
		h = updateHistory(p, h, nextMove, grid)
		shouldRewind = !isPathValid(p, h.path, grid, distance)
		lastMove = ""

		if h.path == "" {
			break
		}
	}
	return p, h
}

func updatePoint(p point, move string, shouldReverse bool) point {
	direction := 1
	if shouldReverse {
		direction = -1
	}

	switch move {
	case "^":
		p.row -= direction
	case "v":
		p.row += direction
	case ">":
		p.col += direction
	case "<":
		p.col -= direction
	}
	return p
}

func updateHistory(p point, h history, move string, grid []string) history {
	h.path += move
	if !isOutsideGrid(p, grid) {
		h.heat += utils.Atoi(string(grid[p.row][p.col]))
	}
	return h
}

func isPathValid(p point, path string, grid []string, distance int) bool {
	debug := true
	if isOutsideGrid(p, grid) {
		log("        isOutsideGrid", debug)
		return false
	}

	if p.row+p.col <= distance {
		log("        isInsideDistance", debug)
		return false
	}

	if getRecentRepeatedMoves(path) > 3 {
		log("        too many recent moves", debug)
		return false
	}

	if loops(p, path) {
		log("        loops", debug)
		return false
	}

	return true
}

func isAlreadyFinalized(p point, finalized map[point]history) bool {
	return p.equals(point{0, 0}) || finalized[p].heat > 0
}

func getRecentRepeatedMoves(history string) int {
	repeats := 0
	if len(history) == 0 {
		return repeats
	}

	direction := ""
	for i := len(history) - 1; i >= 0; i -= 1 {
		char := string(history[i])
		if direction == "" {
			direction = char
		}

		if direction != char {
			break
		}
		repeats++
	}
	return repeats
}

func loops(p point, path string) bool {
	debug := false
	log(fmt.Sprintf("p: %v", p), debug)
	oldP := point{0, 0}
	for _, move := range path {
		if oldP.equals(p) {
			log("path loops", debug)
			return true
		}

		oldP = updatePoint(oldP, string(move), false)
		log(fmt.Sprintf("move: %c => oldP: %v", move, oldP), debug)
	}
	return false
}

func rewind(p point, h history, grid []string) (string, point, history) {
	if h.path == "" {
		return "", p, h
	}

	if !isOutsideGrid(p, grid) {
		h.heat -= utils.Atoi(string(grid[p.row][p.col]))
	}

	lastMove := h.path[len(h.path)-1:]
	p = updatePoint(p, lastMove, true)
	h.path = h.path[:len(h.path)-1]
	return lastMove, p, h
}

func log(message string, shouldLog bool) {
	if shouldLog {
		fmt.Println(message)
	}
}
