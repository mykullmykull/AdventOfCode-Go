package day

import (
	"fmt"
	"goAoc2022/utils"
)

const (
	// directions
	no  = 0
	neg = -1
	pos = 1
	up  = 2
	rt  = 3
	dn  = 4
	lt  = 5

	// sides
	tp = 6
	bm = 7
	ft = 8
	bk = 9
	rs = 10
	ls = 11
)

// rowFrom, colFrom
const (
	first = iota
	rowPos
	rowNeg
	colPos
	colNeg
	last
)

type square struct {
	squareRow int
	squareCol int
}

type point struct {
	row   int
	col   int
	side  int
	toRow int
	toCol int
}

type board struct {
	grid         []string
	walls        []point
	p            point
	instructions string
	isTesting    bool
}

func (p point) turn(inst string) point {
	if inst != "R" && inst != "L" {
		return p
	}
	facing := p.getFacing()
	switch facing {
	case up:
		if inst == "L" {
			return p.changeDirection(lt)
		}
		return p.changeDirection(rt)
	case rt:
		if inst == "L" {
			return p.changeDirection(up)
		}
		return p.changeDirection(dn)
	case dn:
		if inst == "L" {
			return p.changeDirection(rt)
		}
		return p.changeDirection(lt)
	case lt:
		if inst == "L" {
			return p.changeDirection(dn)
		}
		return p.changeDirection(up)
	default:
		panic(fmt.Sprint("invalid facing with toRow ", p.toRow, " and toCol ", p.toCol))
	}
}

func (b board) move(inst string) board {
	if inst == "R" || inst == "L" {
		return b
	}
	dist := 0
	for dist < utils.StrToInt(inst) {
		nextPoint := b.getNextPoint()
		nextChar := b.grid[nextPoint.row][nextPoint.col]
		if nextChar == '.' {
			dist++
			b.p = nextPoint
		}
		if nextChar == '#' {
			return b
		}
	}
	return b
}

func (b board) getNextPoint() point {
	b.p = point{row: b.p.row + b.p.toRow, col: b.p.col + b.p.toCol, side: b.p.side, toRow: b.p.toRow, toCol: b.p.toCol}
	if b.p.row < 0 || b.p.row >= len(b.grid) || b.p.col < 0 || b.p.col >= len(b.grid[0]) || b.grid[b.p.row][b.p.col] == ' ' {
		return b.wrapSides()
	}
	b.p.side = b.getSide(b.p)
	return b.p
}

func (b board) getFacingScore() int {
	switch b.p.getFacing() {
	case rt:
		return 0
	case dn:
		return 1
	case lt:
		return 2
	case up:
		return 3
	}
	panic(fmt.Sprint("invalid facing with toRow ", b.p.toRow, " and toCol ", b.p.toCol))
}
