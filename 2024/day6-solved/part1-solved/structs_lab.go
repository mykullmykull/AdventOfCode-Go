package day

import "fmt"

type lab struct {
	grid  []string
	guard guard
}

func (l lab) findGuard() lab {
	for r, row := range l.grid {
		for c, cell := range row {
			if cell == '^' {
				g := guard{
					row: r,
					col: c,
					dir: up,
				}
				l.guard = g
				return l
			}
		}
	}
	panic("guard not found")
}

func (l lab) moveGuard() lab {
	nextLab, isBlocked := l.getNextStep()
	if isBlocked {
		l.guard = l.guard.turnRight()
		return l
	}
	return nextLab
}

func (l lab) getNextStep() (lab, bool) {
	row := l.guard.row
	col := l.guard.col
	switch l.guard.dir {
	case up:
		row--
	case rt:
		col++
	case dn:
		row++
	case lt:
		col--
	default:
		panic(fmt.Sprintf("Invalid direction %d", l.guard.dir))
	}
	if row >= 0 && row < len(l.grid) && col >= 0 && col < len(l.grid[0]) && l.grid[row][col] == '#' {
		return l, true
	}
	l.guard.row = row
	l.guard.col = col
	return l, false
}

func (l lab) updateGrid(oldGuard guard) lab {
	l.grid[oldGuard.row] = l.grid[oldGuard.row][:oldGuard.col] + "X" + l.grid[oldGuard.row][oldGuard.col+1:]
	newArrow := arrows[l.guard.dir]
	if l.guard.row >= 0 && l.guard.row < len(l.grid) && l.guard.col >= 0 && l.guard.col < len(l.grid[0]) {
		l.grid[l.guard.row] = l.grid[l.guard.row][:l.guard.col] + newArrow + l.grid[l.guard.row][l.guard.col+1:]
	}
	return l
}

func (l lab) isFinished() bool {
	return l.guard.row < 0 || l.guard.row >= len(l.grid) || l.guard.col < 0 || l.guard.col >= len(l.grid[0])
}

func (l lab) countVisitedCells() int {
	count := 0
	for _, row := range l.grid {
		for _, cell := range row {
			if cell == 'X' {
				count++
			}
		}
	}
	return count
}
