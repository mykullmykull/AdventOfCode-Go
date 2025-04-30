package day

import "fmt"

func (b board) wrapSides() point {
	if b.isTesting {
		return b.p.wrapSidesTest(b.getSize(), b.getSquares())
	}
	return b.p.wrapSidesReal(b.getSize(), b.getSquares())
}

func (p point) wrapSidesTest(size int, squares map[int]square) point {
	newSide := -1
	newDirection := -1
	rowFrom := -1
	colFrom := -1

	combined := fmt.Sprint(p.side, p.getFacing())
	switch combined {
	case fmt.Sprint(tp, up):
		newSide = bk
		newDirection = dn
		rowFrom = first
		colFrom = colNeg
	case fmt.Sprint(tp, lt):
		newSide = ls
		newDirection = dn
		rowFrom = first
		colFrom = rowPos
	case fmt.Sprint(tp, rt):
		newSide = rs
		newDirection = lt
		rowFrom = rowNeg
		colFrom = last
	case fmt.Sprint(bk, up):
		newSide = tp
		newDirection = dn
		rowFrom = first
		colFrom = colNeg
	case fmt.Sprint(bk, lt):
		newSide = rs
		newDirection = up
		rowFrom = last
		colFrom = rowNeg
	case fmt.Sprint(bk, dn):
		newSide = bm
		newDirection = up
		rowFrom = last
		colFrom = colNeg
	case fmt.Sprint(ls, up):
		newSide = tp
		newDirection = rt
		rowFrom = colPos
		colFrom = first
	case fmt.Sprint(ls, dn):
		newSide = bm
		newDirection = rt
		rowFrom = colNeg
		colFrom = first
	case fmt.Sprint(ft, rt):
		newSide = rs
		newDirection = dn
		rowFrom = first
		colFrom = rowNeg
	case fmt.Sprint(bm, lt):
		newSide = ls
		newDirection = up
		rowFrom = last
		colFrom = rowNeg
	case fmt.Sprint(bm, dn):
		newSide = bk
		newDirection = up
		rowFrom = last
		colFrom = colNeg
	case fmt.Sprint(rs, up):
		newSide = ft
		newDirection = lt
		rowFrom = colNeg
		colFrom = last
	case fmt.Sprint(rs, rt):
		newSide = tp
		newDirection = lt
		rowFrom = rowNeg
		colFrom = last
	case fmt.Sprint(rs, dn):
		newSide = bk
		newDirection = rt
		rowFrom = first
		colFrom = colNeg
	default:
		panic(fmt.Sprint("invalid side with toRow ", p.toRow, " and toCol ", p.toCol))
	}
	p = p.translate(size, squares[p.side], squares[newSide], rowFrom, colFrom)
	p = p.changeDirection(newDirection)
	p.side = newSide

	return p
}

func (p point) wrapSidesReal(size int, squares map[int]square) point {
	newSide := -1
	newDirection := -1
	rowFrom := -1
	colFrom := -1

	combined := fmt.Sprint(p.side, p.getFacing())
	switch combined {
	case fmt.Sprint(tp, up):
		newSide = bk
		newDirection = rt
		rowFrom = colPos
		colFrom = first
	case fmt.Sprint(tp, lt):
		newSide = ls
		newDirection = rt
		rowFrom = rowNeg
		colFrom = first
	case fmt.Sprint(rs, up):
		newSide = bk
		newDirection = up
		rowFrom = last
		colFrom = colPos
	case fmt.Sprint(rs, rt):
		newSide = bm
		newDirection = lt
		rowFrom = rowNeg
		colFrom = last
	case fmt.Sprint(rs, dn):
		newSide = ft
		newDirection = lt
		rowFrom = colPos
		colFrom = last
	case fmt.Sprint(ft, lt):
		newSide = ls
		newDirection = dn
		rowFrom = first
		colFrom = rowPos
	case fmt.Sprint(ft, rt):
		newSide = rs
		newDirection = up
		rowFrom = last
		colFrom = rowPos
	case fmt.Sprint(ls, up):
		newSide = ft
		newDirection = rt
		rowFrom = colPos
		colFrom = first
	case fmt.Sprint(ls, lt):
		newSide = tp
		newDirection = rt
		rowFrom = rowNeg
		colFrom = first
	case fmt.Sprint(bm, dn):
		newSide = bk
		newDirection = lt
		rowFrom = colPos
		colFrom = last
	case fmt.Sprint(bm, rt):
		newSide = rs
		newDirection = lt
		rowFrom = rowNeg
		colFrom = last
	case fmt.Sprint(bk, lt):
		newSide = tp
		newDirection = dn
		rowFrom = first
		colFrom = rowPos
	case fmt.Sprint(bk, rt):
		newSide = bm
		newDirection = up
		rowFrom = last
		colFrom = rowPos
	case fmt.Sprint(bk, dn):
		newSide = rs
		newDirection = dn
		rowFrom = first
		colFrom = colPos
	default:
		panic(fmt.Sprint("invalid side with toRow ", p.toRow, " and toCol ", p.toCol))
	}

	p = p.translate(size, squares[p.side], squares[newSide], rowFrom, colFrom)
	p = p.changeDirection(newDirection)
	p.side = newSide

	return p
}

func (p point) translate(squareSize int, home square, target square, rowFrom int, colFrom int) point {
	// println("~~translating from", p.toString())
	// fmt.Print("~~with squareSize ", squareSize, " home ", home, " target ", target, " rowFrom ", rowFrom, " colFrom ", colFrom, "\n")

	// normalize
	p.row = p.row - home.squareRow*squareSize
	p.col = p.col - home.squareCol*squareSize

	// get new row
	newRow := -1
	switch rowFrom {
	case first:
		newRow = 0
	case last:
		newRow = squareSize - 1
	case rowPos:
		newRow = p.row
	case rowNeg:
		newRow = squareSize - p.row - 1
	case colPos:
		newRow = p.col
	case colNeg:
		newRow = squareSize - p.col - 1
	}

	// get new col
	newCol := -1
	switch colFrom {
	case first:
		newCol = 0
	case last:
		newCol = squareSize - 1
	case colPos:
		newCol = p.col
	case colNeg:
		newCol = squareSize - p.col - 1
	case rowPos:
		newCol = p.row
	case rowNeg:
		newCol = squareSize - p.row - 1
	}
	// println("~~newRow", newRow, "newCol", newCol)

	// adjust
	p.row = newRow
	p.col = newCol
	p.row = p.row + target.squareRow*squareSize
	p.col = p.col + target.squareCol*squareSize

	// println("~~after adjusting", p.toString())
	return p
}
