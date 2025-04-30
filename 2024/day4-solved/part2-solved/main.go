package day

import "fmt"

func main(grid []string) int {
	foundCount := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			println("looking at r=", r, " c=", c, " with count=", foundCount)

			if isOnEdge(grid, r, c) || grid[r][c] != 'A' {
				continue
			}
			if hasCrosses(grid, r, c) {
				foundCount++
			}
		}
	}
	return foundCount
}

func isOnEdge(grid []string, r int, c int) bool {
	lastRow := len(grid) - 1
	lastCol := len(grid[0]) - 1
	return r == 0 || r == lastRow || c == 0 || c == lastCol
}

func hasCrosses(grid []string, r int, c int) bool {
	fmt.Printf("  looking at r=%d c=%d %c\n", r, c, grid[r][c])
	ul := grid[r-1][c-1]
	ur := grid[r-1][c+1]
	dl := grid[r+1][c-1]
	dr := grid[r+1][c+1]
	fmt.Printf("  ul=%c, ur=%c, dl=%c, dr=%c\n", ul, ur, dl, dr)
	hasCrossUlToDr := (ul == 'M' && dr == 'S') || (ul == 'S' && dr == 'M')
	hasCrossDlToUr := (dl == 'M' && ur == 'S') || (dl == 'S' && ur == 'M')
	println("  hasCrossUlToDr=", hasCrossUlToDr, " hasCrossDlToUr=", hasCrossDlToUr)
	return hasCrossUlToDr && hasCrossDlToUr
}
