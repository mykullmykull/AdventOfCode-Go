package day

import "fmt"

func (p point) getFacing() int {
	if p.isFacing(up) {
		return up
	}
	if p.isFacing(dn) {
		return dn
	}
	if p.isFacing(rt) {
		return rt
	}
	if p.isFacing(lt) {
		return lt
	}
	panic(fmt.Sprint("invalid facing with toRow ", p.toRow, " and toCol ", p.toCol))
}

func (p point) isFacing(direction int) bool {
	shouldBe := point{}.changeDirection(direction)
	return p.toRow == shouldBe.toRow && p.toCol == shouldBe.toCol
}

func (p point) changeDirection(direction int) point {
	switch direction {
	case up:
		p.toRow = neg
		p.toCol = no
	case dn:
		p.toRow = pos
		p.toCol = no
	case rt:
		p.toRow = no
		p.toCol = pos
	case lt:
		p.toRow = no
		p.toCol = neg
	default:
		panic(fmt.Sprint("invalid direction ", direction))
	}
	return p
}
