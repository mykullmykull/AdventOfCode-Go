package day

import "fmt"

const (
	rt = 0
	rd = 1
	dr = 1
	dn = 2
	dl = 3
	ld = 3
	lt = 4
	lu = 5
	ul = 5
	up = 6
	ur = 7
	ru = 7
)

func main(grid []string) int {
	foundCount := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			for d := rt; d <= ru; d++ {
				if isOffEdge(grid, r, c, d) {
					continue
				}
				word := getWord(grid, r, c, d)
				if word == "XMAS" {
					foundCount++
				}
			}
		}
	}
	return foundCount
}

func isOffEdge(grid []string, r int, c int, d int) bool {
	height := len(grid)
	width := len(grid[0])
	switch d {
	case rt:
		return c+3 >= width
	case rd:
		return r+3 >= height || c+3 >= width
	case dn:
		return r+3 >= height
	case dl:
		return r+3 >= height || c-3 < 0
	case lt:
		return c-3 < 0
	case lu:
		return c-3 < 0 || r-3 < 0
	case up:
		return r-3 < 0
	case ur:
		return r-3 < 0 || c+3 >= width
	default:
		panic(fmt.Sprintf("invalid direction %c", d))
	}
}

func getWord(grid []string, r int, c int, d int) string {
	word := ""
	rChange := 0
	cChange := 0
	switch d {
	case rt:
		cChange = 1
	case rd:
		rChange = 1
		cChange = 1
	case dn:
		rChange = 1
	case dl:
		rChange = 1
		cChange = -1
	case lt:
		cChange = -1
	case lu:
		rChange = -1
		cChange = -1
	case up:
		rChange = -1
	case ur:
		rChange = -1
		cChange = 1
	default:
		panic(fmt.Sprintf("invalid direction %c", d))
	}

	for x := 0; x < 4; x++ {
		word += string(grid[r][c])
		r += rChange
		c += cChange
	}
	return word
}
