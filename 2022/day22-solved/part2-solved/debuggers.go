package day

import "fmt"

func (b board) printGrid() {
	copy := make([]string, len(b.grid))
	for row := range b.grid {
		copy[row] = b.grid[row]
	}

	cursor := ""
	switch b.p.getFacing() {
	case up:
		cursor = "^"
	case dn:
		cursor = "v"
	case lt:
		cursor = "<"
	case rt:
		cursor = ">"
	default:
		panic(fmt.Sprint("invalid facing", b.p.getFacing()))
	}

	copy[b.p.row] = copy[b.p.row][:b.p.col] + cursor + copy[b.p.row][b.p.col+1:]
	for _, row := range copy {
		fmt.Println(row, "|")
	}
}

func (b board) printSide() {
	size := b.getSize()
	copy := make([]string, size)
	square := b.getSquares()[b.p.side]
	startRow := square.squareRow * size
	startCol := square.squareCol * size
	endCol := startCol + size

	for x := 0; x < size; x++ {
		copy[x] = b.grid[startRow+x][startCol:endCol]
	}

	cursor := ""
	switch b.p.getFacing() {
	case up:
		cursor = "^"
	case dn:
		cursor = "v"
	case lt:
		cursor = "<"
	case rt:
		cursor = ">"
	default:
		panic(fmt.Sprint("invalid facing", b.p.getFacing()))
	}

	// adjust for the square
	b.p.row -= startRow
	b.p.col -= startCol

	copy[b.p.row] = copy[b.p.row][:b.p.col] + cursor + copy[b.p.row][b.p.col+1:]
	for _, row := range copy {
		fmt.Println(row, "|")
	}
}

func (b board) toString() string {
	str := ""
	for _, row := range b.grid {
		str += fmt.Sprintf("%s|\n", row)
	}

	str += "\n"
	for _, w := range b.walls {
		str += fmt.Sprintf("wall - row: %d, col: %d\n", w.row, w.col)
	}

	str += "\n"
	str += fmt.Sprintf("is testing? %v\n", b.isTesting)
	str += b.p.toString()
	return str
}

func (p point) toString() string {
	return fmt.Sprintf("point - side: %s, row: %d, col: %d, facing: %s", side(p.side), p.row, p.col, facing(p.getFacing()))
}

func facing(n int) string {
	switch n {
	case up:
		return "up"
	case dn:
		return "dn"
	case lt:
		return "lt"
	case rt:
		return "rt"
	default:
		panic(fmt.Sprint("invalid facing enum", n))
	}
}

func side(n int) string {
	switch n {
	case tp:
		return "tp"
	case bm:
		return "bm"
	case ft:
		return "ft"
	case bk:
		return "bk"
	case rs:
		return "rs"
	case ls:
		return "ls"
	default:
		panic(fmt.Sprint("invalid side enum", n))
	}
}
