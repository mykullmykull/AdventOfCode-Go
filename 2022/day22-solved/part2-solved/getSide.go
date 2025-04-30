package day

func (b board) getSide(p point) int {
	if b.isTesting {
		return b.getTestSide(p)
	}
	return b.getRealSide(p)
}

func (b board) getTestSide(p point) int {
	if p.row < 4 {
		return tp
	}
	if p.row < 8 && p.col < 4 {
		return bk
	}
	if p.row < 8 && p.col < 8 {
		return ls
	}
	if p.row < 8 {
		return ft
	}
	if p.col < 12 {
		return bm
	}
	return rs
}

func (b board) getRealSide(p point) int {
	if p.row < 50 && p.col < 100 {
		return tp
	}
	if p.row < 50 {
		return rs
	}
	if p.row < 100 {
		return ft
	}
	if p.row < 150 && p.col < 50 {
		return ls
	}
	if p.row < 150 {
		return bm
	}
	return bk
}

func (b board) getSize() int {
	if b.isTesting {
		return 4
	}
	return 50
}

func (b board) getSquares() map[int]square {
	if b.isTesting {
		return b.getTestSquares()
	}
	return b.getRealSquares()
}

func (b board) getTestSquares() map[int]square {
	return map[int]square{
		tp: {0, 2},
		bk: {1, 0},
		ls: {1, 1},
		ft: {1, 2},
		bm: {2, 2},
		rs: {2, 3},
	}
}

func (b board) getRealSquares() map[int]square {
	return map[int]square{
		tp: {0, 1},
		bk: {3, 0},
		ls: {2, 0},
		ft: {1, 1},
		bm: {2, 1},
		rs: {0, 2},
	}
}
