package day

import (
	"fmt"
	"goAoc2022/utils"
	"math"
)

type board struct {
	grid         []string
	walls        [][]int
	rowWidths    []int
	rowMins      []int
	rowMaxs      []int
	colHeights   []int
	colMins      []int
	colMaxs      []int
	row          int
	col          int
	facingRow    int
	facingCol    int
	instructions string
}

func (b board) popInstruction() (string, board) {
	if b.instructions[0] == 'R' || b.instructions[0] == 'L' {
		inst := string(b.instructions[0])
		b.instructions = b.instructions[1:]
		return inst, b
	}
	inst := ""
	for len(b.instructions) > 0 && b.instructions[0] != 'R' && b.instructions[0] != 'L' {
		inst += string(b.instructions[0])
		b.instructions = b.instructions[1:]
	}
	return inst, b
}

func (b board) turn(inst string) board {
	if inst != "R" && inst != "L" {
		return b
	}
	switch inst {
	case "R":
		if b.facingRow == -1 {
			b.facingRow = 0
			b.facingCol = 1
			break
		}
		if b.facingCol == 1 {
			b.facingRow = 1
			b.facingCol = 0
			break
		}
		if b.facingRow == 1 {
			b.facingRow = 0
			b.facingCol = -1
			break
		}
		if b.facingCol == -1 {
			b.facingRow = -1
			b.facingCol = 0
			break
		}
	case "L":
		if b.facingRow == -1 {
			b.facingRow = 0
			b.facingCol = -1
			break
		}
		if b.facingCol == 1 {
			b.facingRow = -1
			b.facingCol = 0
			break
		}
		if b.facingRow == 1 {
			b.facingRow = 0
			b.facingCol = 1
			break
		}
		if b.facingCol == -1 {
			b.facingRow = 1
			b.facingCol = 0
			break
		}
	}
	return b
}

func (b board) move(inst string) board {
	if inst == "R" || inst == "L" {
		return b
	}
	dist := utils.StrToInt(inst)
	distanceToNearestWall := b.getDistanceToNearestWall()
	dist = utils.UpdateLeast(dist, distanceToNearestWall)
	row := b.row + (b.facingRow * dist)
	col := b.col + (b.facingCol * dist)
	b.row, b.col = b.wrap(row, col)
	return b
}

func (b board) getDistanceToNearestWall() int {
	distance := 1
	row := b.row + b.facingRow
	col := b.col + b.facingCol
	row, col = b.wrap(row, col)
	for !(row == b.row && col == b.col) {
		for _, wCol := range b.walls[row] {
			if wCol == col {
				return distance - 1
			}
		}
		row += b.facingRow
		col += b.facingCol
		row, col = b.wrap(row, col)
		distance++
	}

	return math.MaxInt
}

func (b board) wrap(row int, col int) (int, int) {
	if row == b.row {
		return row, b.wrapCol(row, col)
	}
	if col == b.col {
		return b.wrapRow(row, col), col
	}
	panic(fmt.Sprintf("neither row %d nor col %d are the same as the current row %d or col %d\n", row, col, b.row, b.col))
}

func (b board) wrapRow(row int, col int) int {
	colMin := b.colMins[col]
	colMax := b.colMaxs[col]

	wrapUp := colMin - row
	wrapDn := row - colMax
	if wrapUp > 0 {
		return colMax - (wrapUp % b.colHeights[col]) + 1
	}

	if wrapDn > 0 {
		return colMin + (wrapDn % b.colHeights[col]) - 1
	}

	return row
}

func (b board) wrapCol(row int, col int) int {
	rowMin := b.rowMins[row]
	rowMax := b.rowMaxs[row]

	wrapLt := rowMin - col
	wrapRt := col - rowMax
	if wrapLt > 0 {
		return rowMax - (wrapLt % b.rowWidths[row]) + 1
	}

	if wrapRt > 0 {
		return rowMin + (wrapRt % b.rowWidths[row]) - 1
	}

	return col
}

func (b board) getFacingScore() int {
	if b.facingCol == 1 {
		return 0
	}
	if b.facingRow == 1 {
		return 1
	}
	if b.facingCol == -1 {
		return 2
	}
	if b.facingRow == -1 {
		return 3
	}
	panic(fmt.Sprintf("unknown facing row %d col %d\n", b.facingRow, b.facingCol))
}

func (b board) toString() string {
	str := ""
	for r, row := range b.grid {
		str += fmt.Sprintf("%s - width: %d, min: %d, max: %d\n", row, b.rowWidths[r], b.rowMins[r], b.rowMaxs[r])
	}

	str += "\n"
	for c, _ := range b.grid[0] {
		str += fmt.Sprintf("col: %d, height: %d, min: %d, max: %d\n", c, b.colHeights[c], b.colMins[c], b.colMaxs[c])
	}

	str += "\n"
	for r, row := range b.walls {
		for _, c := range row {
			str += fmt.Sprintf("wall - row: %d, col: %d\n", r, c)
		}
	}
	return str
}
