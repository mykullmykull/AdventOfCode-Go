package day

import "fmt"

type lab struct {
	grid  []string
	guard guard
}

type guard struct {
	row int
	col int
	dir int
}

func (g guard) turnRight() guard {
	g.dir = (g.dir + 1) % 4
	return g
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

func (l lab) doesGuardLoop() bool {
	for {
		// l.printGrid()
		oldGuard := guard{
			row: l.guard.row,
			col: l.guard.col,
			dir: l.guard.dir,
		}
		l = l.moveGuard()
		if l.hasLeftGrid() {
			break
		}
		hasLooped := l.hasLooped()
		if hasLooped {
			return true
		}
		l = l.updateGrid(oldGuard)
	}
	return false
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
	previousDirections := l.grid[oldGuard.row][oldGuard.col : oldGuard.col+1]
	arrow := arrows[oldGuard.dir]
	newDirection := getNewDirection(previousDirections + arrow)
	l.grid[oldGuard.row] = l.grid[oldGuard.row][:oldGuard.col] + newDirection + l.grid[oldGuard.row][oldGuard.col+1:]
	return l
}

func getNewDirection(directions string) string {
	switch directions {
	case "." + arrows[up]:
		return u
	case "." + arrows[dn]:
		return d
	case "." + arrows[lt]:
		return l
	case "." + arrows[rt]:
		return r
	case u + arrows[rt]:
		return ur
	case u + arrows[dn]:
		return ud
	case u + arrows[lt]:
		return ul
	case ur + arrows[dn]:
		return urd
	case ur + arrows[lt]:
		return url
	case urd + arrows[lt]:
		return urdl
	case url + arrows[dn]:
		return urdl
	case ud + arrows[rt]:
		return urd
	case ud + arrows[lt]:
		return udl
	case udl + arrows[rt]:
		return urdl
	case ul + arrows[rt]:
		return url
	case ul + arrows[dn]:
		return udl
	case r + arrows[up]:
		return ur
	case r + arrows[dn]:
		return rd
	case r + arrows[lt]:
		return rl
	case rd + arrows[up]:
		return urd
	case rd + arrows[lt]:
		return rdl
	case rl + arrows[up]:
		return url
	case rl + arrows[dn]:
		return rdl
	case rdl + arrows[up]:
		return urdl
	case d + arrows[up]:
		return ud
	case d + arrows[rt]:
		return rd
	case d + arrows[lt]:
		return dl
	case dl + arrows[up]:
		return ud
	case dl + arrows[rt]:
		return rdl
	case l + arrows[up]:
		return ul
	case l + arrows[rt]:
		return rl
	case l + arrows[dn]:
		return dl
	default:
		panic("invalid direction " + directions)
	}
}

func (l lab) hasLooped() bool {
	prevSymbol := l.grid[l.guard.row][l.guard.col : l.guard.col+1]
	previousDirections := directions[prevSymbol]
	for _, d := range previousDirections {
		if d == l.guard.dir {
			return true
		}
	}
	return false
}

func (l lab) hasLeftGrid() bool {
	return l.guard.row < 0 || l.guard.row >= len(l.grid) || l.guard.col < 0 || l.guard.col >= len(l.grid[0])
}

func (l lab) printGrid() {
	println()
	arrow := arrows[l.guard.dir]
	for _, row := range l.grid {
		if r == fmt.Sprintf("%d", l.guard.row) {
			print(l.grid[l.guard.row][:l.guard.col])
			print(arrow)
			print(l.grid[l.guard.row][l.guard.col+1:] + "\n")
			continue
		}
		println(row)
	}
}
