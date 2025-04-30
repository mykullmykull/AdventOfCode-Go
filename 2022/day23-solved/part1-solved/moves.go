package day

import "fmt"

func (c crater) proposeMoves() crater {
	for y, row := range c.grid {
		for x, cell := range row {
			if cell == '#' {
				c = c.maybeProposeMove(y, x)
			}
		}
	}
	return c
}

func (c crater) maybeProposeMove(y int, x int) crater {
	block := lookAround(c.grid, y, x)
	if countEmpty(block) == 8 {
		return c
	}
	return c.proposeMove(block, y, x)
}

func (c crater) proposeMove(block []string, y int, x int) crater {
	nw := string(block[0][0])
	nn := string(block[0][1])
	ne := string(block[0][2])
	ww := string(block[1][0])
	// center unused
	ee := string(block[1][2])
	sw := string(block[2][0])
	ss := string(block[2][1])
	se := string(block[2][2])

	for _, d := range c.ds {
		isBlocked := false
		switch d {
		case n:
			isBlocked = hasBlocker(nw, nn, ne)
		case e:
			isBlocked = hasBlocker(ne, ee, se)
		case s:
			isBlocked = hasBlocker(sw, ss, se)
		case w:
			isBlocked = hasBlocker(nw, ww, sw)
		default:
			panic("invalid direction " + fmt.Sprint(d))
		}

		if !isBlocked {
			switch d {
			case n:
				block[1] = ww + "^" + ee
				if nn == "O" {
					block[0] = nw + "X" + ne
				} else {
					block[0] = nw + "O" + ne
				}
			case e:
				if ee == "O" {
					block[1] = ww + ">X"
				} else {
					block[1] = ww + ">O"
				}
			case s:
				block[1] = ww + "v" + ee
				if ss == "O" {
					block[2] = sw + "X" + se
				} else {
					block[2] = sw + "O" + se
				}
			case w:
				if ww == "O" {
					block[1] = "X<" + ee
				} else {
					block[1] = "O<" + ee
				}
			default:
				panic("invalid direction " + fmt.Sprint(d))
			}

			c.grid[y-1] = c.grid[y-1][:x-1] + block[0] + c.grid[y-1][x+2:]
			c.grid[y+0] = c.grid[y+0][:x-1] + block[1] + c.grid[y+0][x+2:]
			c.grid[y+1] = c.grid[y+1][:x-1] + block[2] + c.grid[y+1][x+2:]
			break
		}
	}
	return c
}

func hasBlocker(x string, y string, z string) bool {
	return x == "#" || y == "#" || z == "#" ||
		x == "^" || y == "^" || z == "^" ||
		x == ">" || y == ">" || z == ">" ||
		x == "v" || y == "v" || z == "v" ||
		x == "<" || y == "<" || z == "<"
}

func applyMoves(grid []string) []string {
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'O' {
				grid = makeMove(grid, y, x)
			}
			if cell == 'X' {
				grid = undoMoves(grid, y, x)
			}
		}
	}
	return grid
}

func makeMove(grid []string, y int, x int) []string {
	if y > 0 {
		grid = setEmpty(grid, y-1, x, []rune{'v'})
	}
	if x > 0 {
		grid = setEmpty(grid, y, x-1, []rune{'>'})
	}
	if y < len(grid)-1 {
		grid = setEmpty(grid, y+1, x, []rune{'^'})
	}
	if x < len(grid[0])-1 {
		grid = setEmpty(grid, y, x+1, []rune{'<'})
	}

	grid = setElf(grid, y, x, oRunes())
	return grid
}

func undoMoves(grid []string, y int, x int) []string {
	if y > 0 {
		grid = setElf(grid, y-1, x, []rune{'v'})
	}
	if x > 0 {
		grid = setElf(grid, y, x-1, []rune{'>'})
	}
	if y < len(grid)-1 {
		grid = setElf(grid, y+1, x, []rune{'^'})
	}
	if x < len(grid[0])-1 {
		grid = setElf(grid, y, x+1, []rune{'<'})
	}
	grid = setEmpty(grid, y, x, xRunes())
	return grid
}

func setEmpty(grid []string, y int, x int, runes []rune) []string {
	value := rune(grid[y][x])
	for _, r := range runes {
		if value == r {
			grid[y] = grid[y][:x] + "." + grid[y][x+1:]
		}
	}
	return grid
}

func setElf(grid []string, y int, x int, runes []rune) []string {
	value := rune(grid[y][x])
	for _, r := range runes {
		if value == r {
			grid[y] = grid[y][:x] + "#" + grid[y][x+1:]
		}
	}
	return grid
}
