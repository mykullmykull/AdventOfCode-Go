package day

import "fmt"

func (v valley) moveBlizzards() []string {
	newGrid := v.getNewGrid()
	for r, row := range v.grid {
		for c, cell := range row {
			switch cell {
			case '#', '.':
				continue
			case rt:
				newGrid = v.addBlizzard(newGrid, r, c+1, rt)
			case lt:
				newGrid = v.addBlizzard(newGrid, r, c-1, lt)
			case dn:
				newGrid = v.addBlizzard(newGrid, r+1, c, dn)
			case up:
				newGrid = v.addBlizzard(newGrid, r-1, c, up)
			case rl:
				newGrid = v.addBlizzard(newGrid, r, c+1, rt)
				newGrid = v.addBlizzard(newGrid, r, c-1, lt)
			case rd:
				newGrid = v.addBlizzard(newGrid, r, c+1, rt)
				newGrid = v.addBlizzard(newGrid, r+1, c, dn)
			case ru:
				newGrid = v.addBlizzard(newGrid, r, c+1, rt)
				newGrid = v.addBlizzard(newGrid, r-1, c, up)
			case rld:
				newGrid = v.addBlizzard(newGrid, r, c+1, rt)
				newGrid = v.addBlizzard(newGrid, r, c-1, lt)
				newGrid = v.addBlizzard(newGrid, r+1, c, dn)
			case rlu:
				newGrid = v.addBlizzard(newGrid, r, c+1, rt)
				newGrid = v.addBlizzard(newGrid, r, c-1, lt)
				newGrid = v.addBlizzard(newGrid, r-1, c, up)
			case rdu:
				newGrid = v.addBlizzard(newGrid, r, c+1, rt)
				newGrid = v.addBlizzard(newGrid, r+1, c, dn)
				newGrid = v.addBlizzard(newGrid, r-1, c, up)
			case rldu:
				newGrid = v.addBlizzard(newGrid, r, c+1, rt)
				newGrid = v.addBlizzard(newGrid, r, c-1, lt)
				newGrid = v.addBlizzard(newGrid, r+1, c, dn)
				newGrid = v.addBlizzard(newGrid, r-1, c, up)
			case ld:
				newGrid = v.addBlizzard(newGrid, r, c-1, lt)
				newGrid = v.addBlizzard(newGrid, r+1, c, dn)
			case lu:
				newGrid = v.addBlizzard(newGrid, r, c-1, lt)
				newGrid = v.addBlizzard(newGrid, r-1, c, up)
			case ldu:
				newGrid = v.addBlizzard(newGrid, r, c-1, lt)
				newGrid = v.addBlizzard(newGrid, r+1, c, dn)
				newGrid = v.addBlizzard(newGrid, r-1, c, up)
			case du:
				newGrid = v.addBlizzard(newGrid, r+1, c, dn)
				newGrid = v.addBlizzard(newGrid, r-1, c, up)
			default:
				panic("invalid direction " + string(cell))
			}
		}
	}
	return newGrid
}

func (v valley) addBlizzard(grid []string, r int, c int, dir rune) []string {
	combined := combo{dir: dir, current: rune(grid[r][c])}
	if combined.current == '#' {
		r, c = v.wrap(r, c, dir)
		return v.addBlizzard(grid, r, c, dir)
	}
	if combined.current == '.' {
		grid[r] = grid[r][:c] + string(dir) + grid[r][c+1:]
		return grid
	}
	switch combined {
	case combo{rt, lt}:
		dir = rl
	case combo{rt, dn}:
		dir = rd
	case combo{rt, up}:
		dir = ru
	case combo{rt, ld}:
		dir = rld
	case combo{rt, lu}:
		dir = rlu
	case combo{rt, du}:
		dir = rdu
	case combo{rt, ldu}:
		dir = rldu
	case combo{lt, rt}:
		dir = rl
	case combo{lt, dn}:
		dir = ld
	case combo{lt, up}:
		dir = lu
	case combo{lt, rd}:
		dir = rld
	case combo{lt, ru}:
		dir = rlu
	case combo{lt, du}:
		dir = ldu
	case combo{lt, rdu}:
		dir = rldu
	case combo{dn, rt}:
		dir = rd
	case combo{dn, lt}:
		dir = ld
	case combo{dn, up}:
		dir = du
	case combo{dn, rl}:
		dir = rld
	case combo{dn, ru}:
		dir = rdu
	case combo{dn, lu}:
		dir = ldu
	case combo{dn, rlu}:
		dir = rldu
	case combo{up, rt}:
		dir = ru
	case combo{up, lt}:
		dir = lu
	case combo{up, dn}:
		dir = du
	case combo{up, rl}:
		dir = rlu
	case combo{up, rd}:
		dir = rdu
	case combo{up, ld}:
		dir = ldu
	case combo{up, rld}:
		dir = rldu
	default:
		panic(fmt.Sprintf("invalid combo dir: %c and current: %c", combined.dir, combined.current))
	}
	grid[r] = grid[r][:c] + string(dir) + grid[r][c+1:]
	return grid
}

func (v valley) wrap(r int, c int, dir rune) (int, int) {
	if r == 0 && dir == up {
		r = len(v.grid) - 2
	}
	if r == len(v.grid)-1 && dir == dn {
		r = 1
	}
	if c == 0 && dir == lt {
		c = len(v.grid[0]) - 2
	}
	if c == len(v.grid[0])-1 && dir == rt {
		c = 1
	}
	return r, c
}
