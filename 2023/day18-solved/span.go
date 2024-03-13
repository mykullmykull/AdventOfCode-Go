package day18

type span struct {
	minRow int
	maxRow int
	minCol int
	maxCol int
}

func (s span) intersects(other span) bool {
	return other.minCol <= s.maxCol || other.maxCol >= s.minCol
}

func (s span) erases(other span) bool {
	return other.minCol <= s.maxCol || other.maxCol >= s.minCol
}

func (s span) updateWithNewSpan(other span) span {
	// subtracting from left
	if other.minCol == s.minCol {
		s.minCol = other.maxCol
	}

	// adding to right
	if other.minCol == s.maxCol {
		s.maxCol = other.maxCol
	}

	// adding to left
	if other.maxCol == s.minCol {
		s.minCol = other.maxCol
	}

	// subtracting from right
	if other.maxCol == s.maxCol {
		s.maxCol = other.minCol
	}
	return s
}

func (s span) width() int {
	return s.maxCol - s.minCol
}

func (s span) contains(row int, col int) bool {
	return s.minRow <= row && s.maxRow >= row && s.minCol <= col && s.maxCol >= col
}

func (s span) isAbove(row int) bool {
	return s.maxRow < row
}

func (s span) isBelow(row int) bool {
	return s.minRow > row
}

func (s span) isLeft(col int) bool {
	return s.minCol < col
}

func (s span) isRight(col int) bool {
	return s.maxCol > col
}

func (s span) boundsRow(row int) bool {
	return s.minRow <= row && s.maxRow >= row
}

func (s span) boundsCol(col int) bool {
	return s.minCol <= col && s.maxCol >= col
}

func (s span) isVertical() bool {
	return s.minCol == s.maxCol
}

func (s span) isHoritontal() bool {
	return s.minRow == s.maxRow
}
