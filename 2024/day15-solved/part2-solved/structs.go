package day

import "fmt"

type warehouse struct {
	grid  []string
	robot spot
}

type spot struct {
	row int
	col int
}

func (w warehouse) getRobot() warehouse {
	for row := range w.grid {
		for col := range w.grid[row] {
			val := toConst(w.grid[row][col])
			if val == robot {
				w.robot = spot{row, col}
				return w
			}
		}
	}
	panic("No robot found")
}

func (w warehouse) move(d int) warehouse {
	boxesToMove, err := w.getAllBoxesToMove(d)
	if err != nil {
		return w
	}

	return w.moveBoxesAndRobot(boxesToMove, d)
}

func (w warehouse) getAllBoxesToMove(d int) (map[spot]bool, error) {
	boxesToMove := map[spot]bool{}
	pushing := []spot{w.robot}
	for {
		newSpots := w.lookAhead(pushing, d)

		if w.hasWall(newSpots) {
			return map[spot]bool{}, fmt.Errorf("cannot move")
		}
		if w.isOpen(newSpots) {
			break
		}

		newPushing := []spot{}
		for s := range newSpots {
			val := toConst(w.grid[s.row][s.col])
			if val == open {
				continue
			}
			if _, ok := boxesToMove[s]; !ok {
				newPushing = append(newPushing, s)
				boxesToMove[s] = true
			}
		}
		pushing = newPushing
	}
	return boxesToMove, nil
}

func (w warehouse) lookAhead(pushing []spot, d int) map[spot]bool {
	newSpots := map[spot]bool{}
	for _, s := range pushing {
		newRow := s.row
		newCol := s.col
		switch d {
		case up:
			newRow--
		case dn:
			newRow++
		case lt:
			newCol--
		case rt:
			newCol++
		default:
			panic("Invalid direction" + toString(d))
		}
		newSpot := spot{newRow, newCol}
		newSpots[newSpot] = true
	}
	for s := range newSpots {
		val := toConst(w.grid[s.row][s.col])
		if val == boxLt {
			otherHalf := spot{s.row, s.col + 1}
			newSpots[otherHalf] = true
		}
		if val == boxRt {
			otherHalf := spot{s.row, s.col - 1}
			newSpots[otherHalf] = true
		}
	}
	return newSpots
}

func (w warehouse) hasWall(spots map[spot]bool) bool {
	for s := range spots {
		if toConst(w.grid[s.row][s.col]) == wall {
			return true
		}
	}
	return false
}

func (w warehouse) isOpen(spots map[spot]bool) bool {
	for s := range spots {
		if toConst(w.grid[s.row][s.col]) != open {
			return false
		}
	}
	return true
}

func (w warehouse) moveBoxesAndRobot(boxesToMove map[spot]bool, d int) warehouse {
	// copy grid
	newGrid := make([]string, len(w.grid))
	copy(newGrid, w.grid)

	// translate d into row and col changes
	rowChange, colChange := translateIntoChanges(d)

	// move each box
	for b := range boxesToMove {
		val := toConst(w.grid[b.row][b.col])
		newRow := b.row + rowChange
		newCol := b.col + colChange
		newGrid[newRow] = newGrid[newRow][:newCol] + toString(val) + newGrid[newRow][newCol+1:]
	}

	// move robot
	newRobot := spot{w.robot.row + rowChange, w.robot.col + colChange}
	newGrid[newRobot.row] = newGrid[newRobot.row][:newRobot.col] + toString(robot) + newGrid[newRobot.row][newRobot.col+1:]
	newGrid[w.robot.row] = newGrid[w.robot.row][:w.robot.col] + toString(open) + newGrid[w.robot.row][w.robot.col+1:]

	// erase any left over half boxes
	newGrid = eraseHalfBoxes(newGrid)

	return warehouse{grid: newGrid, robot: newRobot}
}

func translateIntoChanges(d int) (int, int) {
	rowChange := 0
	colChange := 0
	switch d {
	case up:
		rowChange = -1
	case dn:
		rowChange = 1
	case lt:
		colChange = -1
	case rt:
		colChange = 1
	default:
		panic("Invalid direction" + toString(d))
	}
	return rowChange, colChange
}

func eraseHalfBoxes(grid []string) []string {
	for r, row := range grid {
		for c := range row {
			val := toConst(grid[r][c])
			if val != boxLt && val != boxRt {
				continue
			}
			lVal := toConst(grid[r][c-1])
			rVal := toConst(grid[r][c+1])
			if (val == boxLt && rVal != boxRt) || (val == boxRt && lVal != boxLt) {
				grid[r] = grid[r][:c] + toString(open) + grid[r][c+1:]
			}
		}
	}
	return grid
}
